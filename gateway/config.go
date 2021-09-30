package gateway

import "github.com/501miles/go-tiny/model"

const Default_Config_Name = "config.yml"

type Config struct {
	BaseConfig BaseConfig `yaml:"base"`
	ConsulConfig model.ConsulConfig `yaml:"consul"`
}

type BaseConfig struct {
	Ip         string `yaml:"ip"`
	Port       uint16 `yaml:"port"`
	InstanceId uint32 `yaml:"instance_id"`
	ServerId   uint32 `yaml:"server_id"`
}