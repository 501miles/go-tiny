package server

import (
	"fmt"
	"github.com/501miles/logger"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

type BaseService struct {
	name   string
	sid    uint32
	ip     string
	port   uint16
	secure bool
	config interface{}
	consulClient *api.Client
}

func (b *BaseService) Name() string {
	return b.name
}

func (b *BaseService) SID() uint32 {
	return b.sid
}

func (b *BaseService) IP() string {
	return b.ip
}

func (b *BaseService) Port() uint16 {
	return b.port
}

func (b *BaseService) IsSecure() bool {
	return b.secure
}

func (b *BaseService) Init() error {
	b.port = 12306
	b.name = "test1"
	b.sid = 112233
	b.ip = "192.168.1.233"

	//TODO 从config文件读取配置并赋值
	config := api.DefaultConfig()
	config.Address = "www.evan0.xyz:8501"
	config.Token = "7f85db13-c45f-f619-3acc-756d2d9af9cf"
	c, err := api.NewClient(config)
	if err != nil {
		return err
	}
	b.consulClient = c
	return nil
}

func (b *BaseService) ServeCallback() error {
	go func() {
		r := gin.Default()
		// /head/login/1
		urlPath := fmt.Sprintf("/health/%s/%d", b.Name(), b.SID())
		logger.Info(urlPath)
		// 健康检测接口，其实只要是 200 就认为成功了
		r.GET(urlPath, func(c *gin.Context) {
			c.JSON(200,nil)
		})
		err := r.Run(":8090")
		if err != nil {
			logger.Error(err)
		}
	}()
	return nil
}

func (b *BaseService) Start() error {
	err := b.ServeCallback()
	if err != nil {
		return err
	}
	return b.RegisterService()
}

func (b *BaseService) RegisterService() error {
	registration := new(api.AgentServiceRegistration)
	registration.ID = fmt.Sprintf("%s-%d", b.Name(), b.SID())
	registration.Name = registration.ID
	registration.Port = int(b.Port())
	registration.Address = b.IP()

	check := new(api.AgentServiceCheck)
	schema := "http"
	if b.IsSecure() {
		schema = "https"
	}
	check.HTTP = fmt.Sprintf("%s://%s:%d/actuator/health", schema, registration.Address, registration.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "20s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check
	err := b.consulClient.Agent().ServiceRegister(registration)
	if err != nil {
		return err
	}
	return nil
}

func (b *BaseService) DeregisterService() {
	logger.Info("DeregisterService")
	err := b.consulClient.Agent().ServiceDeregister(fmt.Sprintf("%s-%d", b.Name(), b.SID()))
	if err != nil {
		logger.Error(err)
	}
}

func (b *BaseService) Shutdown() error {
	return nil
}

func (b *BaseService) Ping() uint8 {
	return 1
}
