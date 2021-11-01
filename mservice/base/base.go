package base

import (
	"context"
	"fmt"
	"github.com/501miles/go-tiny/rpc/message"
	"github.com/501miles/go-tiny/tool/logx"
	"github.com/501miles/go-tiny/tool/time_tool"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	logger "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net"
)

type BaseService struct {
	name         string
	instanceId   uint32
	serverId     uint32
	ip           string
	port         uint16
	secure       bool
	config       Config
	consulClient *api.Client
}

func (b *BaseService) Name() string {
	return b.name
}

func (b *BaseService) SID() uint32 {
	return b.instanceId
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
	logx.Init()
	serviceConfig := Config{}
	yamlFile, err := ioutil.ReadFile(Default_Config_Name)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &serviceConfig)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	b.config = serviceConfig
	b.port = serviceConfig.BaseConfig.Port
	b.name = serviceConfig.BaseConfig.Name
	b.instanceId = serviceConfig.BaseConfig.InstanceId
	b.ip = serviceConfig.BaseConfig.Ip

	config := api.DefaultConfig()
	config.Address = serviceConfig.ConsulConfig.IP
	if serviceConfig.ConsulConfig.Port > 0 {
		config.Address = fmt.Sprintf("%s:%d",config.Address, serviceConfig.ConsulConfig.Port)
	}
	config.Token = serviceConfig.ConsulConfig.Token
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
			c.JSON(200, nil)
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
		return fmt.Errorf("ServeCallback start faild: %v", err)
	}
	//err = b.RegisterService()
	//if err != nil {
	//	return fmt.Errorf("RegisterService start faild: %v", err)
	//
	//}
	err = b.StartRPCServer()
	return err
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

func (b *BaseService) StartRPCServer() error {
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	message.RegisterGatewayServiceServer(grpcServer, b)
	err = grpcServer.Serve(lis)
	if err != nil {
		return fmt.Errorf("failed to serve: %v", err)
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

func (b *BaseService) Version() string {
	return "1.0"
}


func (b *BaseService) RequestService(ctx context.Context, in *message.GatewayMsg) (*message.ResMsg, error) {
	logger.Info("调用RequestService")
	reqMsg := &message.ReqMsg{
		RequestData: in.RequestData,
		UserId:      in.UserId,
	}
	bs := b.ProcessRPCRequest(reqMsg)
	return &message.ResMsg{
		MsgId:        in.MsgId,
		T:            time_tool.NowTimeUnix13(),
		ResponseData: bs,
	}, nil
}

func (b *BaseService) ProcessRPCRequest(msg *message.ReqMsg) []byte {
	var bs []byte
	return bs
}