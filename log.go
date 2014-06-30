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
	this.doLogWriter(text, LEVEL_ALL)
}
func (this Log) Debug(text string) {
	this.doLogWriter(text, LEVEL_DEBUG)
}
func (this Log) Warn(text string) {
	this.doLogWriter(text, LEVEL_WARN)
}
func (this Log) Error(text string) {
	this.doLogWriter(text, LEVEL_ERROR)
}
func (this Log) Fatal(text string) {
	this.doLogWriter(text, LEVEL_FATAL)
}
func (this Log) doLogWriter(text string, level int) bool {
	isTrue := false
	if level < this.defaultLevel {
		return isTrue
	}
	this.logWriter.Write(text, level)
	isTrue = true
	return isTrue
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
	log.defaultLevel = LEVEL_WARN
}

//~log.go************************************************
