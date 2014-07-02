package log4go_test

import (
	"log4go"
)

func ExampleWriter() {
	log4go.All("Hello This is Test Log4go.All()")
	log4go.Debug("Hello This is Test Log4go.Debug()")
	log4go.Warn("Hello This is Test Log4go.Warn()")
	log4go.Error("Hello This is Test Log4go.Error()")
	log4go.Fatal("Hello This is Test Log4go.Fatal()")
}
