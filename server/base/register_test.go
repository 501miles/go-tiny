package base

import (
	logger "github.com/sirupsen/logrus"
	"testing"
)

func Test1(t *testing.T)  {
	s1 := new(BaseService)
	err := s1.Init()
	if err != nil {
		logger.Error(err)
	}
	err = s1.Start()
	if err != nil {
		logger.Error(err)
	}
	select {}
}
