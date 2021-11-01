package base

import "github.com/501miles/go-tiny/rpc/message"

type MService interface {
	Name() string
	SID() uint32
	IP() string
	Port() uint16
	Init() error
	RegisterService() error
	DeregisterService()
	ServeCallback() error
	Start() error
	Shutdown() error
	Ping() uint8
	IsSecure() bool
	Version() string
	StartRPCServer() error
	ProcessRPCRequest(msg *message.ReqMsg) []byte
}
