package WeHub

import (
	"NewGateServer/database"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/silenceper/wechat/user"
	"log"
	"sync"
	"time"
)

type UserInfo struct {
	UnionID         string                 `json:"union_id" bson:"union_id"`
	CTX             map[string]interface{} `json:"ctx" bson:"ctx"`
	LastMessageTime int64                  `json:"last_message_time" bson:"last_message_time"`
}

type Info struct {
	Subscribe      int32   `json:"subscribe"`
	Nickname       string  `json:"nickname"`
	Sex            int32   `json:"sex"`
	City           string  `json:"city"`
	Country        string  `json:"country"`
	Province       string  `json:"province"`
	Language       string  `json:"language"`
	Headimgurl     string  `json:"headimgurl"`
	SubscribeTime  int32   `json:"subscribe_time"`
	UnionID        string  `json:"unionid"`
	Remark         string  `json:"remark"`
	GroupID        int32   `json:"groupid"`
	TagidList      []int32 `json:"tagid_list"`
	SubscribeScene string  `json:"subscribe_scene"`
	UpdateTime     int64   `json:"update_time" bson:"update_time"`
}

func (i *Info) CopyFromWeInfo(info *user.Info) {
	i.Subscribe = info.Subscribe
	i.Nickname = info.Nickname
	i.Sex = info.Sex
	i.City = info.City
	i.Country = info.Country
	i.Province = info.Province
	i.Language = info.Language
	i.Headimgurl = info.Headimgurl
	i.SubscribeTime = info.SubscribeTime
	i.UnionID = info.UnionID
	i.Remark = info.Remark
	i.GroupID = info.GroupID
	i.TagidList = info.TagidList
	i.SubscribeScene = info.SubscribeScene
	i.UpdateTime = time.Now().Unix()
}

func GetUsersInfoTest(ctx iris.Context) {

	var q struct {
		Appid   string `json:"appid"`
		Refresh bool   `json:"refresh"`
	}
	err := ctx.ReadQuery(&q)
	if err != nil {
		ctx.JSON(iris.Map{"status": 1, "msg": "param error"})
		return
	}
	// 从参数中获取appid

	appid := q.Appid
	// 是否更新
	refresh := q.Refresh

	if appid == "" {
		ctx.JSON(iris.Map{"status": 1, "msg": "appid empty",})
		return
	}

	var getKey = func(appid string) string {
		return fmt.Sprintf("%s_%s", database.DBUsersInfo, appid)
	}

	ds := database.NewSessionStore()
	defer ds.Close()
	con := ds.C(getKey(appid))

	if !refresh {
		var uinfo []user.Info
		con.Find(nil).All(&uinfo)
		ctx.JSON(iris.Map{"status": 0, "msg": "ok", "res": uinfo})
		return
	}

	var umlist []string
	if s, exists := hub.Servers[appid]; exists {
		log.Printf("开始更新%s的数据\n", s.cfg.NickName)
		umlist, err = s.Wc.GetUser().ListAllUserOpenIDs()
		if err != nil {
			log.Printf("更新失败%s\n", s.cfg.NickName)
			ctx.JSON(iris.Map{"status": 1, "msg": err.Error()})
			return
		}
		if len(umlist) == 0 {
			ctx.JSON(iris.Map{"status": 1, "msg": "no user"})
			return
		}
	}

	var wg sync.WaitGroup
	wg.Add(len(umlist))

	var ResChan = make(chan interface{}, len(umlist))
	for _, k := range umlist {
		fmt.Println("用户", k)
		go func(Res chan interface{}, wg *sync.WaitGroup) {

			uinfo, err := hub.Servers[appid].Wc.GetUser().GetUserInfo(k)
			if err != nil {
				log.Printf("获取%s信息失败,%s", k, err)
				return
			}
			fmt.Println("输入用户数据", uinfo)
			ResChan <- uinfo
			wg.Done()
		}(ResChan, &wg)
	}

	var tempinfos []interface{}

	go func() {
		wg.Wait()
		close(ResChan)
	}()

	for {
		select {
		case res, ok := <-ResChan:
			if ok {
				con.Insert(res)
				tempinfos = append(tempinfos, res)
			} else {
				ctx.JSON(iris.Map{"status": 0, "msg": "ok", "res": tempinfos})
				return
			}
		}
	}
}
