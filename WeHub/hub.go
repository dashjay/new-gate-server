package WeHub

import (
	"NewGateServer/configs"
	"NewGateServer/database"
	"NewGateServer/utils"
	"errors"
	"fmt"
	"github.com/night-codes/mgo-ai"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"sync"
	"time"
)

var hub *Hub

func init() {
	hub = NewHub()
	// 从数据库中载入数据 并且注入Hub中
	ds := database.NewSessionStore()
	defer ds.Close()
	var cfgs []*Config
	err := ds.C(database.DBServer).Find(nil).All(&cfgs)
	if err != nil {
		panic(err)
	}
	for _, cfg := range cfgs {

		hub.Servers[cfg.APPID] = new(Servers)
		hub.Servers[cfg.APPID].cfg = cfg

		// 注册路由
		hub.Mux[cfg.Suffix] = cfg.APPID
		// 注册ID
		rcache := hub.NewRedisOption()

		hub.Servers[cfg.APPID].Wc = wechat.NewWechat(&wechat.Config{
			AppID:          cfg.APPID,
			AppSecret:      cfg.APPSecret,
			Token:          cfg.Suffix,
			EncodingAESKey: cfg.EncodingAESKey,
			Cache:          rcache,
		})

		log.Printf("Service Create, NickName: %s,Suffix: %s,HandlerAddr: %s\n", cfg.NickName, cfg.Suffix, cfg.HandlerAddr)
	}
	// end载入数据中已存的公众号信息bu

	// 将数据库中的用户数据全部导入应用。
	var uinfos []UserInfo
	err = ds.C(database.DBUsersCTX).Find(nil).All(&uinfos)
	if err != nil {
		panic(err)
	}
	// 把用户信息给存进数据库里
	for _, uinfo := range uinfos {
		hub.UINFO.Set(uinfo.UnionID, uinfo)
	}
	go hub.UserInfoSync()
}

// Hub 整个微信服务的中心
type Hub struct {
	Servers map[string]*Servers  // 存储不同的服务节点
	Mux     map[string]string    // 服务路由mux
	UINFO   *utils.ConCurrentMap // 用户信息存储
	RedisDB int                  // redisdb number
	cnum    int                  // http clint的池子中的对象数量
	Client  *sync.Pool           // http对象池
}

// NewHub 初始化为6个服务
func NewHub() *Hub {
	return &Hub{
		Servers: make(map[string]*Servers, 6),
		RedisDB: 0,
		Mux:     make(map[string]string, 6),
		cnum:    0,
		UINFO:   utils.NewConCurrentMap(),
		Client: &sync.Pool{New: func() interface{} {
			// 方便后期获取Client的个数查看使用情况
			hub.cnum++
			return &http.Client{Timeout: 3}
		}},
	}
}

// CreateServer 创建一个服务信息
func (h *Hub) CreateServer(cfg *Config) error {

	// 检测cfg缺少参数
	if !cfg.Valid() {
		return errors.New("config invalid")
	}
	// 存在该服务
	if _, exists := h.Servers[cfg.APPID]; exists {
		return errors.New(fmt.Sprintf("config exists, appid: %s", cfg.APPID))
	}

	// 创建
	h.Servers[cfg.APPID] = new(Servers)
	// 赋值config
	hub.Servers[cfg.APPID].cfg = cfg

	ds := database.NewSessionStore()
	defer ds.Close()

	// 注册ID
	ai.Connect(ds.C("counters"))
	hub.Servers[cfg.APPID].cfg.ID = ai.Next("servers")

	err := ds.C("servers").Insert(&hub.Servers[cfg.APPID].cfg)
	if err != nil {
		return err
	}
	// 相当于路由
	hub.Mux[cfg.Suffix] = cfg.APPID

	rcache := h.NewRedisOption()

	hub.Servers[cfg.APPID].Wc = wechat.NewWechat(&wechat.Config{
		AppID:          cfg.APPID,
		AppSecret:      cfg.APPSecret,
		Token:          cfg.Suffix,
		EncodingAESKey: cfg.EncodingAESKey,
		Cache:          rcache,
	})

	log.Printf("Service Create NickName: %s,Suffix: %s,HandlerAddr: %s\n", cfg.NickName, cfg.Suffix, cfg.HandlerAddr)
	return nil
}

// UpdateServer 更新服务信息，将接受到的服务信息更新到对应服务中
func (h *Hub) UpdateServer(cfg *Config) error {

	// 检测cfg缺少参数
	if !cfg.Valid() {
		return errors.New("config invalid")
	}
	// 不存在报告
	if _, exists := h.Servers[cfg.APPID]; !exists {
		return errors.New(fmt.Sprintf("config not exists, appid: %s", cfg.APPID))
	}

	// 获取目前的ID
	currentID := h.Servers[cfg.APPID].cfg.ID
	// Update
	h.Servers[cfg.APPID].cfg = cfg
	// 保持ID不变
	h.Servers[cfg.APPID].cfg.ID = currentID

	ds := database.NewSessionStore()
	defer ds.Close()

	// 更新
	err := ds.C("servers").Update(bson.M{"appid": cfg.APPID}, &h.Servers[cfg.APPID].cfg)
	if err != nil {
		return err
	}

	rcache := h.NewRedisOption()

	hub.Servers[cfg.APPID].Wc = wechat.NewWechat(&wechat.Config{
		AppID:          cfg.APPID,
		AppSecret:      cfg.APPSecret,
		Token:          cfg.Suffix,
		EncodingAESKey: cfg.EncodingAESKey,
		Cache:          rcache,
	})

	log.Printf("Service Update Success NickName: %s,Suffix: %s,HandlerAddr: %s\n", cfg.NickName, cfg.Suffix, cfg.HandlerAddr)
	return nil
}

// NewRedisOption 创建一个Redis的配置，并且++
func (h *Hub) NewRedisOption() *cache.Redis {
	defer func() {
		h.RedisDB++
	}()
	return cache.NewRedis(&cache.RedisOpts{
		Host:        configs.C.RedisHost,
		Database:    h.RedisDB,
		MaxIdle:     20,
		MaxActive:   200,
		IdleTimeout: 24,
	})
}

// UserInfoSync 在设定的的时间内，周期性的将所有系统中存在的
// 用户信息同步到数据库中
func (h *Hub) UserInfoSync() {
	tick := time.NewTicker(time.Hour * 12)
	for {
		select {
		case <-tick.C:
			{
				// 创建会话
				ds := database.NewSessionStore()
				con := ds.C(database.DBUsersCTX)
				log.Println("开始同步用户消息至数据库")
				total := 0
				// 遍历全部对象
				for k, v := range h.UINFO.Items() {
					con.Upsert(bson.M{"union_id": k}, v)
					total++
				}

				log.Printf("总共同步了%d个用户信息", total)
			}
		}
	}
}
