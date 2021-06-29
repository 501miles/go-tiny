package time_tool

import (
	"math"
	"time"
)

// NowTimeUnix13 unix时间戳13位，精确到毫秒
func NowTimeUnix13() int64 {
	return time.Now().UnixNano() / 1e6
}

// NowTimeUnix10 unix时间戳10位，精确到秒
func NowTimeUnix10() int64 {
	return time.Now().Unix()
}

// NowTimeUnixWithLength unix时间戳len位，len范围[10,19]
func NowTimeUnixWithLength(len int) int64 {
	if len < 10 || len > 19 {
		return time.Now().UnixNano()
	}
	return time.Now().UnixNano() / int64(math.Pow10(19-len))
}

// NowTimeFormattedStr 格式化的时间字符串2006-01-02 15:04:05
func NowTimeFormattedStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
