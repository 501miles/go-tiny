package rpc

import (
	"github.com/501miles/go-tiny/rpc/message"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type ServiceRPCClient struct {
	Conn *grpc.ClientConn
	Client *message.GatewayServiceClient
	ServiceName string
	ServiceInstanceId uint32
	RequestMessage proto.Message
	ResponseMessage proto.Message
}