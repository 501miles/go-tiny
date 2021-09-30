package model

type MysqlConfig struct {
	Ip string
	Port int
	Name string
	User string
	Pwd string
}

type RedisConfig struct {
	Ip string
	Port int
	Pwd string
	DbNum int
}

type MongoConfig struct {
	Ip string
	Port int
	Name string
	User string
	Pwd string
}

type ConsulConfig struct {
	IP    string `yaml:"ip"`
	Port  uint16 `yaml:"port"`
	Token string `yaml:"token"`
}
