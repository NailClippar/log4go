/********************************************************
*														*
*	title:log.go										*
*	author:NailClippar									*
*	date:2014-6-30										*
*	env:Ubuntu 14.04 LTS x64 with Golang 1.2.1			*
*														*
*********************************************************
 */
package log4go

//log level
const (
	LEVEL_ALL = iota
	LEVEL_DEBUG
	LEVEL_WARN
	LEVEL_ERROR
	LEVEL_FATAL
	LEVEL_OFF
)

//static 方法
func All(text string) {
	log.All(text)
}
func Debug(text string) {
	log.Debug(text)
}
func Warn(text string) {
	log.Warn(text)
}
func Error(text string) {
	log.Error(text)
}
func Fatal(text string) {
	log.Fatal(text)
}

func SetDisplayLevel(level int) {
	log.defaultLevel = level
}

//Log类型
type Log struct {
	logWriter    LogWriter
	defaultLevel int
}

func (this Log) All(text string) {
}
func (this Log) Debug(text string) {
}
func (this Log) Warn(text string) {
}
func (this Log) Error(text string) {
}
func (this Log) Fatal(text string) {
}
func (this Log) Off(text string) {

}

var log *Log

//Logger接口
//用于对外调用
type Logger interface {
	All(text string)
	Debug(text string)
	Warn(text string)
	Error(text string)
	Fatal(text string)
	Off(text string)
}

//LogWriter类型
//用于声明Writer
type LogWriter interface {
	Write(text string, level int)
}

//init method

func init() {
	log = new(Log)
	log.logWriter = NewConsoleWriter()
}

//~log.go************************************************
