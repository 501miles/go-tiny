package gateway

const Default_Config_Name = "config.yml"

type Config struct {
	Ip   string `yaml:"ip"`
	Port string `yaml:"port"`
}
