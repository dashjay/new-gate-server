package WeHub

import (
	"NewGateServer/configs"
	"NewGateServer/database"
	"errors"
	"fmt"
	"github.com/night-codes/mgo-ai"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"sync"
)

var hub *Hub

func init() {
	hub = NewHub()

	// 从数据库中载入数据 并且注入Hub中
	ds := database.NewSessionStore()
	defer ds.Close()
	con := ds.C("servers")
	var cfgs []*Config
	err := con.Find(nil).All(&cfgs)
	if err != nil {
		panic(err)
	}
	for _, cfg := range cfgs {
		fmt.Println(cfg)
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
}

type Hub struct {
	Servers map[string]*Servers
	Mux     map[string]string
	RedisDB int
	cnum    int
	Client  *sync.Pool
}

// NewHub 初始化为6个服务
func NewHub() *Hub {
	return &Hub{
		Servers: make(map[string]*Servers, 6),
		RedisDB: 0,
		Mux:     make(map[string]string, 6),
		cnum:    0,
		Client: &sync.Pool{New: func() interface{} {
			// 方便后期获取Client的个数查看使用情况
			hub.cnum++
			return &http.Client{Timeout: 3}
		}},
	}
}

func (h *Hub) CreateServer(cfg *Config) error {

	// 检测cfg缺少参数
	if !cfg.Valid() {
		return errors.New("config invalid")
	}
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
