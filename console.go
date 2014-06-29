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
	"runtime"
)

//初始化一些环境参数
const (
	NewLine = "\n"
)

//画刷
type Brush func(string) string

func NewBrush(color string) Brush {
	preText := "\033[1;"
	resetText := "\033[0m\n"
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
	NewBrush("0")}

//ConsoleWriter

type ConsoleWriter struct {
}

func (cw *ConsoleWriter) Write(text string) {
	fmt.Printf(text)
}
func NewConsoleWriter() LogWriter {
	var logWriter LogWriter = new(ConsoleWriter)
	return logWriter
}
func init() {
	fmt.Printf("HEHE")
	var GOOS string = runtime.GOOS
	fmt.Printf(GOOS + "\n")
}

//~console.go********************************************
