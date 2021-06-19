package server

type MService interface {
	Name() string
	SID() uint32
	IP() string
	Port() string
	Start() error
	Shutdown() error
	Ping() uint8
}