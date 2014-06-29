package log4go

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	for level,color := range colors{
		fmt.Printf("%d"+color("Hello"),level)
	}
}