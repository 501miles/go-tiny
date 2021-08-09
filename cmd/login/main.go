package main

import (
	"github.com/501miles/go-tiny/cmd/login/service"
	"github.com/501miles/logger"
)

func main() {
	s := new(service.LoginService)
	err := s.Init()
	if err != nil {
		logger.Fatal(err)
	}
	err = s.Start()
	if err != nil {
		logger.Fatal(err)
	}
}