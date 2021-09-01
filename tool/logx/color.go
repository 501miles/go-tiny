// +build !debug disable_color

package logx

func Blue(str string) string    { return "\x1b[0;34m" + str + "\x1b[0m" }
func Yellow(str string) string  { return "\x1b[0;33m" + str + "\x1b[0m" }
func Green(str string) string   { return "\x1b[0;32m" + str + "\x1b[0m" }
func Magenta(str string) string { return "\x1b[0;35m" + str + "\x1b[0m" }
func Cyan(str string) string    { return "\x1b[0;36m" + str + "\x1b[0m" }
func Gray(str string) string    { return "\x1b[0;37m" + str + "\x1b[0m" }
func White(str string) string   { return "\x1b[0;30m" + str + "\x1b[0m" }
func Red(str string) string     { return "\x1b[0;31m" + str + "\x1b[0m" }

func DebugMsg(str string) string { return colors[6](str) }
func TraceMsg(str string) string { return colors[1](str) }

func RedMsg(str string) string    { return colors[3](str) }
func BlueMsg(str string) string   { return colors[5](str) }
func YellowMsg(str string) string { return colors[4](str) }
func Blue2Msg(str string) string  { return colors[2](str) }

type brush func(string) string

func newBrush(color string) brush {
	pre := "\033["
	reset := "\033[0m"
	return func(text string) string {
		return pre + color + "m" + text + reset
	}
}

var colors = []brush{
	newBrush("1;41"), // Emergency          红色底
	newBrush("1;35"), // Alert              紫色
	newBrush("1;34"), // Critical           蓝色
	newBrush("1;31"), // Error              红色
	newBrush("1;33"), // Warn               黄色
	newBrush("1;36"), // Informational      天蓝色
	newBrush("1;32"), // Debug              绿色
	newBrush("1;32"), // Trace              绿色
}
