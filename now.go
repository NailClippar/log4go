/********************************************************
*														*
*	title:now.go										*
*	author:NailClippar									*
*	date:2014-6-30										*
*	env:Ubuntu 14.04 LTS x64 with Golang 1.2.1			*
*														*
*********************************************************
 */
//~now.go************************************************
package log4go

import (
	"time"
)

const DefaultTimeFormat = "2006-01-02 15:04:05"

func NowStr() string {
	return time.Now().Format(DefaultTimeFormat)
}
