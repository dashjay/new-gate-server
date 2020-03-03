package configs

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Configs struct {
	RedisHost string      `toml:"RedisHost"`
	Port      int         `toml:"Port"`
	MongoDB   MongoConfig `toml:"MongoDB"`
}

var C Configs

type MongoConfig struct {
	DBName           string
	DBHost           string
	DBUser           string
	DBPassword       string
	MongoDBPoolLimit int
}

func init() {
	_, err := toml.DecodeFile("./configs/configs.toml", &C)
	log.Println(C)
	if err != nil {
		panic(err)
	}
}
