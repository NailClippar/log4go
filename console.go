/********************************************************
*														*
*	title:console.go 									*
*	author:NailClippar									*
*	date:2014-6-29										*
*	env:Ubuntu 14.04 LTS x64 with Golang 1.2.1			*
*														*
*********************************************************
 */

/*

*/

package log4go

import (
	"fmt"
	"os"
	"runtime"
)

//初始化一些环境参数
var NewLine = "\n"

//画刷
type Brush func(string) string

func NewBrush(color string) Brush {
	preText := "\033[1;"
	resetText := "\033[0m"
	return func(logText string) string {
		return preText + color + "m" + logText + resetText
	}
}

var colors = []Brush{
	NewBrush("30"), //All black
	NewBrush("34"), //Debug cyan
	NewBrush("32"), //INFO green
	NewBrush("33"), //Warn yellow
	NewBrush("31"), //Error red
	NewBrush("35"), //FATAL purple
	NewBrush("0")}  //OFF Close

//levelDetail
var levelDetail = []string{
	"A",
	"D",
	"W",
	"E",
	"F",
	"O"}

//ConsoleWriter

type ConsoleWriter struct {
}

func (cw *ConsoleWriter) Write(text string, level int) {
	textStr := colors[level](text)
	timeStr := NowStr()
	levelStr := "[" + colors[level](levelDetail[level]) + "]"
	fmt.Fprintf(os.Stdout, "%s %s:%s\n", timeStr, levelStr, textStr)
}
func NewConsoleWriter() LogWriter {
	var logWriter LogWriter = new(ConsoleWriter)
	return logWriter
}
func init() {
	//env check
	if runtime.GOOS == "windows" {
		NewLine = "\r\n"
	}
}

//~console.go********************************************
