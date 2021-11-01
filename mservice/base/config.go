package base

import (
	"github.com/501miles/go-tiny/model"
)

const Default_Config_Name = "config.yml"

type Config struct {
	BaseConfig BaseConfig `yaml:"base_config"`
	MysqlConfig model.MysqlConfig `yaml:"mysql_config"`
	RedisConfig model.RedisConfig `yaml:"redis_config"`
	MongoConfig model.MongoConfig `yaml:"mongo_config"`
	ConsulConfig model.ConsulConfig `yaml:"consul_config"`
}

type BaseConfig struct {
	Name       string `yaml:"name"`
	Ip         string `yaml:"ip"`
	Port       uint16 `yaml:"port"`
	InstanceId uint32 `yaml:"instance_id"`
	ServerId   uint32 `yaml:"server_id"`
}



func readConfig() {

}
