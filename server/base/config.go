package base

import (
	"github.com/501miles/go-tiny/model"
)

type Config struct {
	BaseConfig BaseConfig
	MysqlConfig model.MysqlConfig
	RedisConfig model.RedisConfig
	MongoConfig model.MongoConfig
}

type BaseConfig struct {
	Name string `yaml:"name"`
	Ip string `yaml:"ip"`
	Port int `yaml:"port"`
	InstanceId int `yaml:"instance_id"`
	ServerId int `yaml:"server_id"`
}



func readConfig() {

}
