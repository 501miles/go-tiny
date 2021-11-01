package query

import (
	"github.com/501miles/go-tiny/rpc/message"
	"github.com/501miles/go-tiny/server/base"
)

type Service struct {
	base.BaseService
}

func (s *Service) ProcessRPCRequest(msg *message.ReqMsg) []byte {
	return []byte("query service.")
}