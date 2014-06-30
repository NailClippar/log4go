package log4go

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	for level, color := range colors {
		fmt.Printf("%d"+color("Hello")+"\n", level)
	}
}
func TestConsoleWriter(t *testing.T) {
	NewConsoleWriter().Write("This is Test ConsoleWriter", LEVEL_DEBUG)
}
