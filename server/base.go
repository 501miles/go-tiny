package server

type BaseService struct {
	name   string
	sid    uint32
	ip     string
	port   string
	config interface{}
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

func (b *BaseService) Port() string {
	return b.port
}

func (b *BaseService) Start() error {
	return nil
}

func (b *BaseService) Shutdown() error {
	return nil
}

func (b *BaseService) Ping() uint8 {
	return 1
}
