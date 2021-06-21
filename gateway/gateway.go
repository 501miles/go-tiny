package gateway

import "sync"

type GateWay struct {
	IP       string
	Port     string
	Services sync.Map
}
