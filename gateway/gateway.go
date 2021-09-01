package gateway

import (
	"errors"
	"fmt"
	"github.com/501miles/go-tiny/tool/logx"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

type GateWay struct {
	IP         string
	Port       string
	InstanceId int
	ServerId   int
	Services   sync.Map
}

var gateWayConfig Config

func NewGateway() *GateWay {
	return &GateWay{
		IP:         "",
		Port:       "",
		InstanceId: 0,
		ServerId:   0,
		Services:   sync.Map{},
	}
}

func (g *GateWay) Init() error {
	logx.Init()
	gateWayConfig = Config{}
	yamlFile, err := ioutil.ReadFile(Default_Config_Name)
	if err != nil {
		return errors.New(fmt.Sprintf("yamlFile.Get err   #%v ", err))
	}

	err = yaml.Unmarshal(yamlFile, &gateWayConfig)
	if err != nil {
		return errors.New(fmt.Sprintf("Unmarshal: %v ", err))
	}
	g.IP = gateWayConfig.Ip
	g.Port = gateWayConfig.Port
	g.InstanceId = gateWayConfig.InstanceId
	g.ServerId = gateWayConfig.ServerId
	return nil
}

func (g *GateWay) Run() {
	logger.Infof("gateway instance_id: %d, server_id %d, is running...", g.InstanceId, g.ServerId)
	r := gin.Default()
	err := r.Run(fmt.Sprintf("%s: %s", gateWayConfig.Ip, gateWayConfig.Port))
	if err != nil {
		logger.Fatalf("gateway instance_id: %d, server_id %d, fail to start", g.InstanceId, g.ServerId)
	}
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

//开始从consul获取注册在相同serverId下的微服务
func monitorMService() {

}


