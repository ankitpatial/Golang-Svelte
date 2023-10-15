package infra

import (
	"fmt"
	"github.com/fatih/color"
)

var green = color.New(color.FgGreen)

func LogInfo(msg string, arg ...interface{}) {
	if len(arg) > 0 {
		msg = fmt.Sprintf(msg, arg...)
	}
	_, _ = green.Println(msg)
}
