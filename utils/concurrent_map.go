package utils

import (
	"github.com/orcaman/concurrent-map"
)

type ConCurrentMap struct {
	m cmap.ConcurrentMap
}

func NewConCurrentMap() *ConCurrentMap {
	conMap := cmap.New()
	return &ConCurrentMap{m: conMap}
}

func (c *ConCurrentMap) Get(key string) (interface{}, bool) {
	v, ok := c.m.Get(key)
	return v, ok
}
func (c *ConCurrentMap) Set(key string, value interface{}) {
	c.m.Set(key, value)
}

func (c *ConCurrentMap) Del(key interface{}) {
	c.m.Remove(key.(string))
}

func (c *ConCurrentMap) Keys() []string {
	return c.m.Keys()
}

func (c *ConCurrentMap) Items() map[string]interface{} {
	return c.m.Items()
}

var UpsertFunc = func(exist bool, valueInMap interface{}, newValue interface{}) interface{} {
	if exist {
		return newValue
	} else {
		return valueInMap
	}
}

func (c *ConCurrentMap) Upsert(key string, v interface{}) {
	c.m.Upsert(key, v, UpsertFunc)
}
