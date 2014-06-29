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
}
func Debug(text string) {
}
func Warn(text string) {
}
func Error(text string) {

}
func FATAL(text string) {

}

func SetDisplayLevel(level int) {

}

//Log类型
type Log struct {
	logWriter  LogWriter
	defaultLvl int
}

func (this *Log) All(text string) {
}
func (this *Log) Debug(text string) {
}
func (this *Log) Warn(text string) {
}
func (this *Log) Error(text string) {
}
func (this *Log) FATAL(text string) {
}

var log *Log

//Logger接口
//用于对外调用
type Logger interface {
	All(text string)
	Debug(text string)
	Warn(text string)
	Error(text string)
	FATAL(text string)
}

//LogWriter类型
//用于声明Writer
type LogWriter interface {
	Write(text string)
}

//init method

func init() {
	log = new(Log)
	log.logWriter = NewConsoleWriter()
}
