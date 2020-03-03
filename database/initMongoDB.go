package database

import (
	. "NewGateServer/configs"
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session

// InitMongoDB 初始化一个MongoDB会话，并持有该链接
func init() {

	dialInfo := &mgo.DialInfo{
		Addrs:     []string{C.MongoDB.DBHost},
		Direct:    false,
		Timeout:   time.Second * 60,
		Username:  C.MongoDB.DBUser,
		Password:  C.MongoDB.DBPassword,
		PoolLimit: C.MongoDB.MongoDBPoolLimit,
	}

	var err error
	session, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
}

type SessionStore struct {
	session *mgo.Session
}

func (d *SessionStore) C(name string) *mgo.Collection {
	return d.session.DB(C.MongoDB.DBName).C(name)
}

//为每一HTTP请求创建新的DataStore对象
func NewSessionStore() *SessionStore {

	ds := &SessionStore{
		session: session.Copy(),
	}
	return ds
}

func (d *SessionStore) Close() {
	d.session.Close()
}

func GetErrNotFound() error {
	return mgo.ErrNotFound
}

func Over() {
	session.Close()
}
