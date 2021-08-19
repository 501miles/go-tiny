package main

import (
	"github.com/501miles/go-tiny/gateway"
	"github.com/501miles/logger"
)

func main() {
	err := gateway.Init()
	if err != nil {
		logger.Fatal(err)
	}
	gateway.RunApiGateway()
}
