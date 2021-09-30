package gateway

import (
	context2 "context"
	"errors"
	"fmt"
	"github.com/501miles/go-tiny/gateway/rpc"
	"github.com/501miles/go-tiny/model"
	"github.com/501miles/go-tiny/rpc/message"
	"github.com/501miles/go-tiny/tool/gen_id/snowflake"
	"github.com/501miles/go-tiny/tool/logx"
	"github.com/501miles/go-tiny/tool/time_tool"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/hashicorp/consul/api"
	logger "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

type GateWay struct {
	IP            string
	Port          uint16
	InstanceId    uint32
	ServerId      uint32
	slock         sync.RWMutex
	ClientManager *rpc.ClientManager
	ServiceMap    map[string][]*model.Service
	consulClient  *api.Client
}

var gateWayConfig Config

func NewGateway() *GateWay {
	return &GateWay{
		ServiceMap: map[string][]*model.Service{},
		ClientManager: rpc.NewClientManager(),
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
	g.IP = gateWayConfig.BaseConfig.Ip
	g.Port = gateWayConfig.BaseConfig.Port
	g.InstanceId = gateWayConfig.BaseConfig.InstanceId
	g.ServerId = gateWayConfig.BaseConfig.ServerId
	snowflake.Init(int32(g.InstanceId))
	go g.monitorMService()
	return nil
}

func (g *GateWay) Run() {
	logger.Infof("gateway instance_id: %d, server_id %d, is running...", g.InstanceId, g.ServerId)
	r := gin.Default()
	//TODO 动态注册路由?
	param := message.QueryParam{
		Key:   "student",
		Param: map[string]string{},
	}
	bs, _ := proto.Marshal(&param)
	r.GET("/v1/query", func(context *gin.Context) {
		logger.Info(g.ClientManager.GetRPCClientByName("query"))
		res, err := (*g.ClientManager.GetRPCClientByName("query").Client).RequestService(context2.Background(), &message.GatewayMsg{
			Version:     1,
			T:           time_tool.NowTimeUnix13(),
			MsgId:       g.genMsgId(),
			RequestData: bs,
			UserId:      110,
		})
		if err != nil {
			logger.Error(err)
		}
		logger.Info(res)
		context.Writer.WriteString("ok")
	})
	err := r.Run(fmt.Sprintf("%s:%d", gateWayConfig.BaseConfig.Ip, gateWayConfig.BaseConfig.Port))
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
func(g *GateWay) monitorMService() {
	//

	g.ClientManager.AddMService(model.Service{
		Name:       "query",
		Address:    "127.0.0.1",
		Port:       9999,
		Protocol:   "grpc",
		ServerId:   100,
		InstanceId: 123,
	})
}

func (g *GateWay) registerMService(s *model.Service) {
	g.slock.Lock()
	defer g.slock.Unlock()
	list, ok := g.ServiceMap[s.Name]
	if !ok {
		list = make([]*model.Service, 2)
	}
	list = append(list, s)
	g.ServiceMap[s.Name] = list
}

func (g *GateWay) unregisterMService() {
	g.slock.Lock()
	defer g.slock.Unlock()

}

func callRPC() {

}

func (g *GateWay) genMsgId() int64 {
	return snowflake.GenInt64()
}

