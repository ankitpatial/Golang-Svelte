package rt

import (
	"fmt"
	"roofix/config"
	"runtime"
	"strings"
)

type Stack []uintptr

// CallerInfo will be used with runtime.Caller(1)
func CallerInfo(_ uintptr, file string, line int, ok bool) string {
	if !ok {
		return ""
	}

	path := fmt.Sprintf("%s:%d", file, line)
	idx := strings.Index(path, fmt.Sprintf("/%s/", config.AppName)) + 1
	return path[idx:]
}

func Callers() *Stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st Stack = pcs[0:n]
	return &st
}
