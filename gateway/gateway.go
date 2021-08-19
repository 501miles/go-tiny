package gateway

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

type GateWay struct {
	IP       string
	Port     string
	Services sync.Map
}

var gateWayConfig Config

func Init() error {
	gateWayConfig = Config{}
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		return errors.New(fmt.Sprintf("yamlFile.Get err   #%v ", err))
	}

	err = yaml.Unmarshal(yamlFile, &gateWayConfig)
	if err != nil {
		return errors.New(fmt.Sprintf("Unmarshal: %v ", err))
	}
	return nil
}

func RunApiGateway() {
	r := gin.Default()
	r.Run(fmt.Sprintf("%s: %s", gateWayConfig.Ip, gateWayConfig.Port))
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With, Token, User")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		c.Writer.Header().Set("x-frame-options", "SAMEORIGIN")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
