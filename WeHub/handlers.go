package WeHub

import (
	"NewGateServer/database"
	"bytes"
	"errors"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/silenceper/wechat/message"
	"github.com/silenceper/wechat/miniprogram"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type Config struct {
	ID             uint64 `json:"id" bson:"id"`             // 一个自增的ID, 表示我们创建了多少服务
	Name           string `json:"name" bson:"name"`         // 每个服务的英文名
	NickName       string `json:"nickname" bson:"nickname"` // 每个服务的中文名
	APPID          string `json:"appid" bson:"appid"`
	APPSecret      string `json:"app_secret" bson:"app_secret"`
	Suffix         string `json:"suffix" bson:"suffix"` // 服务后缀
	HandlerAddr    string `json:"handler_addr"`         // 服务的addr
	EncodingAESKey string `json:"encoding_aes_key"`
	Type           string `json:"type" ` // minip 小程序 official 公众号
}

// Valid 检查是否合法
func (c *Config) Valid() bool {
	return c.Name != "" && c.APPID != "" && c.APPSecret != "" && c.Suffix != "" && c.HandlerAddr != "" && c.Type != ""
}

var (
	httpreg = regexp.MustCompile(`http*`)
	grpcreg = regexp.MustCompile(`grpc*`)
)

// Create 创建一个程序放入Hub中
func Create(ctx iris.Context) {
	var Errchan = make(chan error)

	go func(e chan error) {
		config := new(Config)
		err := ctx.ReadJSON(config)
		if err != nil {
			e <- err
			return
		}
		// 引用
		err = hub.CreateServer(config)
		if err != nil {
			e <- err
		}
		e <- nil
		return
	}(Errchan)

	err := <-Errchan
	if err != nil {
		ctx.JSON(iris.Map{"status": 1, "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"status": 0, "msg": "create success"})
}

// Del 从hub中删除一个程序
func Del(ctx iris.Context) {
	appid := ctx.FormValue("appid")
	//appid := ctx.Params().Get("appid")
	if appid == "" {
		ctx.JSON(iris.Map{"status": 1, "msg": "appid empty"})
		return
	}
	if a, exists := hub.Servers[appid]; exists {
		defer log.Printf("Del Server, appid: %s, appname: %s", a.cfg.APPID, a.cfg.NickName)
		// 取到后缀
		var suffix = a.cfg.Suffix
		// 删除
		delete(hub.Servers, appid)
		// 删除路由
		delete(hub.Mux, suffix)

		// 真.从数据库中删除
		ds := database.NewSessionStore()
		defer ds.Close()

		err := ds.C(database.DBServer).Remove(bson.M{"appid": a.cfg.APPID})
		if err != nil {
			ctx.JSON(iris.Map{"status": 1, "msg": err.Error()})
			return
		}

		ctx.JSON(iris.Map{"status": 0, "msg": "delete success"})
		return
	} else {
		ctx.JSON(iris.Map{"status": 1, "msg": "appid not exists"})
	}
}

// Monitor 检查所有程序
func Monitor(ctx iris.Context) {
	var res []iris.Map
	for _, v := range hub.Servers {
		var total = 0.0
		var ave = 0.0
		if len(v.ResponseTime) > 0 {
			for _, f := range v.ResponseTime {
				total += f
			}
			ave = total / float64(len(v.ResponseTime))
		}

		_, exists := hub.Mux[v.cfg.Suffix]

		res = append(res, iris.Map{
			"ave_response_time": ave,
			"cfg":               v.cfg,
			"status":            exists,
		})
	}
	_, err := ctx.JSON(&res)
	if err != nil {
		log.Panic(err)
	}

}

func Stop(ctx iris.Context) {
	appid := ctx.FormValue("appid")
	if appid == "" {
		ctx.JSON(iris.Map{"status": 1, "msg": "appid empty"})
		return
	}
	if s, exists := hub.Servers[appid]; exists {
		if appid, exists2 := hub.Mux[s.cfg.Suffix]; exists2 {
			log.Printf("Stop Server, appid: %s, appname: %s\n", appid, s.cfg.NickName)
			delete(hub.Mux, s.cfg.Suffix)
			ctx.JSON(iris.Map{"status": 0, "msg": "stop success!"})
			return
		} else {
			ctx.JSON(iris.Map{"status": 1, "msg": "it stop already!"})
			return
		}
	} else {
		ctx.JSON(iris.Map{"status": 1, "msg": fmt.Sprintf("server not exists, appid: %s", appid)})
	}
}

func Start(ctx iris.Context) {
	appid := ctx.FormValue("appid")
	if appid == "" {
		ctx.JSON(iris.Map{"status": 1, "msg": "appid empty"})
		return
	}
	if s, exists := hub.Servers[appid]; exists {
		if _, exists2 := hub.Mux[s.cfg.Suffix]; exists2 {
			ctx.JSON(iris.Map{"status": 1, "msg": "is start already!"})
			return
		} else {
			// 加入路由
			hub.Mux[s.cfg.Suffix] = appid
			log.Printf("Start Server, appid: %s, appname: %s", s.cfg.APPID, s.cfg.NickName)
			ctx.JSON(iris.Map{"status": 0, "msg": "start success"})
			return
		}
	} else {
		ctx.JSON(iris.Map{"status": 1, "msg": fmt.Sprintf("server not exists, appid: %s", appid)})
	}
}

func Update(ctx iris.Context) {
	var Errchan = make(chan error)

	go func(e chan error) {
		config := new(Config)
		err := ctx.ReadJSON(config)
		if err != nil {
			e <- err
			return
		}
		// 引用
		err = hub.UpdateServer(config)
		if err != nil {
			e <- err
		}
		e <- nil
		return
	}(Errchan)

	err := <-Errchan
	if err != nil {
		ctx.JSON(iris.Map{"status": 1, "msg": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"status": 0, "msg": "update success"})
}

// MainHandler 主服务
func MainHandler(ctx iris.Context) {
	path := strings.Split(ctx.Path(), "/")
	suffix := path[len(path)-1]

	if appid, exists := hub.Mux[suffix]; exists {
		var start = time.Now()

		switch hub.Servers[appid].cfg.Type {

		case "minip":
			var ResChan = make(chan interface{})
			var ErrChan = make(chan error)
			go func(reschan chan<- interface{}, errchan chan<- error) {
				// 取到minip的上下文
				m := hub.Servers[appid].Wc.GetMiniProgram()
				// 获取用户传参Code
				code := ctx.Params().Get("code")
				if code == "" {
					errchan <- errors.New("code empty")
					return
				}
				// 进行Code2Session
				res, err := m.Code2Session(code)
				if err != nil {
					errchan <- err
					return
				}
				if res.ErrCode != 0 {
					errchan <- errors.New(res.ErrMsg)
					return
				}
				var r ResCode2Session
				// 从Code2Session 讲获得的相关信息Copy到用户信息中
				r.CopyFromCode2Session(&res)

				// 给信息打上AppID 方便分离应用
				r.AppID = appid

				ds := database.NewSessionStore()
				defer ds.Close()
				con := ds.C(database.MiniProgramUsers)

				// 查询用户使用情况
				var origin ResCode2Session
				// 从数据库中查找该用户的信息

				// 指定该用户针对该应用的的selector
				var selector = bson.M{"appid": appid, "unionid": r.UnionID}

				if err := con.Find(selector).One(&origin); err != nil {
					if err.Error() == database.GetErrNotFound().Error() {

						// 第一次使用该APP
						err := con.Insert(&r)
						if err != nil {
							errchan <- err
							return
						}
						// 成功
						reschan <- r
						return
					} else {
						errchan <- err
					}
				}
				// 原来就有数据

				// 更新信息
				_ = con.Update(selector, r)
				ResChan <- r
				return
			}(ResChan, ErrChan)

			select {
			case e := <-ErrChan:
				ctx.JSON(iris.Map{"status": 1, "msg": e.Error()})
				return
			case r := <-ResChan:
				end := time.Now()
				// 计算时间
				hub.Servers[appid].ResponseTime = append(hub.Servers[appid].ResponseTime, end.Sub(start).Seconds())
				// 超过50的扔掉
				if len(hub.Servers[appid].ResponseTime) > 50 {
					hub.Servers[appid].ResponseTime = hub.Servers[appid].ResponseTime[1:]
				}
				ctx.JSON(iris.Map{"status": 0, "msg": "code2session success", "object": r})
				return
			case <-time.After(2 * time.Second):
				ctx.JSON(iris.Map{"status": 1, "msg": "timeout"})
			}
		case "official":
			server := hub.Servers[appid].Wc.GetServer(ctx.Request(), ctx.ResponseWriter())
			server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

				if httpreg.MatchString(hub.Servers[appid].cfg.HandlerAddr) {
					var bm = BaseMessage{Message: &msg}
					bytebm, err := bm.MarshalJSON()
					if err != nil {
						return NewTextMessage(err.Error())
					}
					client := hub.Client.Get().(*http.Client)
					defer hub.Client.Put(client)
					req, err := http.NewRequest("POST", hub.Servers[appid].cfg.HandlerAddr, bytes.NewReader(bytebm))
					if err != nil {
						return NewTextMessage(err.Error())
					}
					resp, err := client.Do(req)
					if err != nil {
						return NewTextMessage(err.Error())
					}
					defer resp.Body.Close()
					content, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						return NewTextMessage(err.Error())
					}
					var reply BaseReply
					err = reply.UnmarshalJSON(content)
					if err != nil {
						return NewTextMessage(err.Error())
					}

					return reply.Reply
				}
				return NewTextMessage("未指定HandlerAddr")
			})
		default:
			panic(hub.Servers[appid])
			return
		}

	} else {
		log.Printf("suffix: %s not exists", suffix)
	}
}

func NewTextMessage(text string) *message.Reply {
	return &message.Reply{MsgType: message.MsgTypeText, MsgData: message.NewText(text)}
}

// ResCode2Session 小程序快速Code2Session
type ResCode2Session struct {
	AppID      string `json:"appid" bson:"appid"`
	OpenID     string `json:"openid" bson:"openid"`       // 用户唯一标识
	SessionKey string `json:"session_key" bson:"session"` // 会话密钥
	UnionID    string `json:"unionid" bson:"unionid"`     // 用户在开放平台的唯一标识符，在满足UnionID下发条件的情况下会返回
}

func (r *ResCode2Session) Reset() {
	r.SessionKey = ""
	r.UnionID = ""
	r.OpenID = ""
	return
}

func (r *ResCode2Session) CopyFromCode2Session(session *miniprogram.ResCode2Session) {
	r.Reset()
	r.OpenID = session.OpenID
	r.UnionID = session.UnionID
	r.SessionKey = session.SessionKey
	return
}
