package log4go

import (
	"fmt"
	"testing"
)

func TestNowStr(t *testing.T) {
	fmt.Println(NowStr() + "时间测试字符串")
}
