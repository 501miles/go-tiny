package timeline

import (
	"github.com/501miles/go-tiny/tool/gen_id/snowflake"
	"github.com/501miles/logger"
	"testing"
	"time"
)

func Test(t *testing.T) {
	snowflake.Init(1)
	timeline, err := CreateTimeline(1, 1*time.Second)
	if err != nil {
		logger.Error(err)
	}

	for i := 0; i < 1000; i++ {
		logger.Info(i)
		func(i2 int) {
			timeline.Register(time.Now().Add(time.Duration(i)*time.Second), func() {
				logger.Info(i2)
			})
		}(i)

	}
	timeline.Start()
}
