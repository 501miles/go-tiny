package model

type Dog struct {
	Name string
	Kind string
	Age  int
}

// Service 外部用
type Service struct {
	Name       string
	Address    string
	Port       uint16
	Protocol   string
	ServerId   uint32
	InstanceId uint32
}