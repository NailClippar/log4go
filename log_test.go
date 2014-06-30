package log4go

import "testing"

func TestWrite(t *testing.T) {
	log.logWriter.Write("呵呵\n", LEVEL_FATAL)
}

func TestLog(t *testing.T) {
	log.All("Hello This is Test Log.All()")
	log.Debug("Hello This is Test Log.Debug()")
	log.Warn("Hello This is Test Log.Warn()")
	log.Error("Hello This is Test Log.Error()")
	log.Fatal("Hello This is Test Log.Fatal()")
}
