package log4go

import "testing"

func TestWrite(t *testing.T) {
	log.logWriter.Write("呵呵\n", LEVEL_FATAL)
}
