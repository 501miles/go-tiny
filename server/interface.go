package server

type MService interface {
	Name() string
	SID() uint32
	IP() string
	Port() string
	Init() error
	RegisterService()
	Start() error
	Shutdown() error
	Ping() uint8
}
