package base

type Config struct {
	BaseConfig BaseConfig
	MysqlConfig MysqlConfig
	RedisConfig RedisConfig
	MongoConfig MongoConfig
}

type BaseConfig struct {
	Name string
	Ip string
	Port int
	InstanceId int
	ServerId int
}

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

func readConfig() {

}
