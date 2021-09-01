package logx

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
)

// MyFormatter 自定义的日志格式化器
type MyFormatter struct {
	IsColor  bool
	FullPath bool
}

func (m *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	var newLog string
	//HasCaller()为true才会有调用信息
	if entry.HasCaller() {
		newLog = fmt.Sprintf("[%s] [%s] [%s] %s\n",
			m.colorTime(entry), m.colorLevel(entry), m.colorCaller(entry), m.colorMessage(entry))
	} else {
		newLog = fmt.Sprintf("[%s] [%s] %s\n", m.colorTime(entry), m.colorLevel(entry), m.colorMessage(entry))
	}

	b.WriteString(newLog)
	return b.Bytes(), nil
}

func (m *MyFormatter) colorCaller(entry *logrus.Entry) string {
	file := entry.Caller.File
	line := entry.Caller.Line
	str := ""
	if !m.FullPath {
		file = path.Base(entry.Caller.File)
		str = fmt.Sprintf(" %s:%d ", file, line)
	} else {
		str = fmt.Sprintf("%s:%d", file, line)
	}
	if m.IsColor {
		if m.FullPath {
			return Blue2Msg(fmt.Sprintf(" %s ", str))
		}
		return Blue2Msg(str)
	} else {
		return str
	}
}

func (m *MyFormatter) colorTime(entry *logrus.Entry) string {
	str := entry.Time.Format("2006-01-02 15:04:05")
	if m.IsColor {
		switch entry.Level {
		case logrus.ErrorLevel:
			return RedMsg(str)
		case logrus.WarnLevel:
			return YellowMsg(str)
		case logrus.DebugLevel:
			return DebugMsg(str)
		case logrus.TraceLevel:
			return TraceMsg(str)
		default:
			return BlueMsg(str)
		}
	} else {
		return str
	}
}

func (m *MyFormatter) colorLevel(entry *logrus.Entry) string {
	str := levelToString(entry.Level)
	if m.IsColor {
		switch entry.Level {
		case logrus.ErrorLevel:
			return RedMsg(str)
		case logrus.WarnLevel:
			return YellowMsg(str)
		case logrus.DebugLevel:
			return DebugMsg(str)
		case logrus.TraceLevel:
			return TraceMsg(str)

		default:
			return BlueMsg(str)
		}
	} else {
		return str
	}

}

func (m *MyFormatter) colorMessage(entry *logrus.Entry) string {
	str := entry.Message
	if m.IsColor {
		switch entry.Level {
		case logrus.ErrorLevel:
			return RedMsg(str)
		case logrus.WarnLevel:
			return YellowMsg(str)
		case logrus.DebugLevel:
			return DebugMsg(str)
		case logrus.TraceLevel:
			return TraceMsg(str)
		default:
			return BlueMsg(str)
		}
	} else {
		return str
	}
}

func levelToString(level logrus.Level) string {
	switch level {
	case logrus.TraceLevel:
		return "TRAC"
	case logrus.DebugLevel:
		return "DEBG"
	case logrus.InfoLevel:
		return "INFO"
	case logrus.WarnLevel:
		return "WARN"
	case logrus.ErrorLevel:
		return "EROR"
	case logrus.FatalLevel:
		return "FTAL"
	case logrus.PanicLevel:
		return "PANI"

	default:
		return "UKNW"
	}
}
