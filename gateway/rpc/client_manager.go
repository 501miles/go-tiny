package rpc

import (
	"errors"
	"fmt"
	"github.com/501miles/go-tiny/model"
	"github.com/501miles/go-tiny/rpc/message"
	"google.golang.org/grpc"
	"math/rand"
	"sync"
	"time"
)

type ClientManager struct {
	lock sync.RWMutex
	serviceRPCClientDict map[uint32]*ServiceRPCClient
	serviceRPCClientNameDict map[string][]*ServiceRPCClient
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		serviceRPCClientDict: map[uint32]*ServiceRPCClient{},
		serviceRPCClientNameDict: map[string][]*ServiceRPCClient{},
	}
}

func (cm * ClientManager) AddMService(s model.Service) error {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	if s.InstanceId <= 0 {
		return errors.New("error service < 0")
	}
	sc := &ServiceRPCClient{
		ServiceName:       s.Name,
		ServiceInstanceId: s.InstanceId,
	}

	address := fmt.Sprintf("%s:%d", s.Address, s.Port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("faild to connect: %v", err)
	}
	client := message.NewGatewayServiceClient(conn)
	sc.Client = &client
	cm.serviceRPCClientDict[s.InstanceId] = sc
	if list, ok := cm.serviceRPCClientNameDict[s.Name]; ok {
		list = append(list, sc)
		cm.serviceRPCClientNameDict[s.Name] = list
	}else{
		cm.serviceRPCClientNameDict[s.Name] = []*ServiceRPCClient{sc}
	}

	return nil
}

func (cm * ClientManager) GetRPCClientByInstanceId(id uint32) *ServiceRPCClient {
	cm.lock.RLock()
	defer cm.lock.RUnlock()
	return cm.serviceRPCClientDict[id]
}

func (cm *ClientManager) GetRPCClientByName(name string) *ServiceRPCClient {
	//获取策略:随机
	cm.lock.RLock()
	defer cm.lock.RUnlock()
	if sc, ok := cm.serviceRPCClientNameDict[name]; ok {
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)
		return sc[r.Intn(len(sc))]
	}
	return nil
}