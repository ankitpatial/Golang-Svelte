package log

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"roofix/config"
	lErr "roofix/pkg/util/log/errors"
	"roofix/pkg/util/rt"
	"runtime"
	"time"
)

var (
	green  = color.New(color.FgGreen)
	blue   = color.New(color.FgBlue)
	yellow = color.New(color.FgYellow)
	red    = color.New(color.FgRed)
	isDev  bool
)

func init() {
	log.SetFlags(0)
	isDev = config.IsDevEnv()
}

func ts() string {
	if isDev {
		return time.Now().Format("2006-01-02T15:04:05")
	}

	// on server aws logs itself has time stamps
	return ""
}

func Info(msg string, arg ...any) {
	caller := rt.CallerInfo(runtime.Caller(1))

	a := make([]any, 0, 3+len(arg))
	a = append(a, ts(), blue.Sprint("INFO"), caller)
	a = append(a, arg...)

	log.Printf("%s %v %s - "+msg, a...)
}

func InfoBullet(msg string, arg ...any) {
	caller := rt.CallerInfo(runtime.Caller(1))

	a := make([]any, 0, 3+len(arg))
	a = append(a, ts(), green.Sprint("INFO"), caller)
	a = append(a, arg...)

	log.Printf("%s %v %s ✦ "+msg, a...)
}

func Warn(msg string, arg ...interface{}) {
	caller := rt.CallerInfo(runtime.Caller(1))
	a := make([]any, 0, 3+len(arg))
	a = append(a, ts(), yellow.Sprint("WARN"), caller)
	a = append(a, arg...)

	log.Printf("%s %v %s - ⚠️ "+msg, a...)
}

func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s: %w", message, err)
}

// Error will log err to std output as well as dump it to storage
func Error(err error) {
	if err == nil {
		return
	}

	caller := rt.CallerInfo(runtime.Caller(1))

	er := lErr.New(err)
	errMsg := er.Error()
	stack := string(er.Stack())

	a := make([]any, 0, 4)
	a = append(a, ts(), red.Sprint("ERROR"), errMsg, stack)

	log.Printf("%s %v %s \n ** Call Trace ** \n%s\n", a...)

	dump(caller, errMsg, stack)
}

// PrintError to std out
func PrintError(err error) {
	if err == nil {
		return
	}

	er := lErr.New(err)
	errMsg := er.Error()
	stack := string(er.Stack())

	a := make([]any, 0, 4)
	a = append(a, ts(), red.Sprint("ERROR"), errMsg, stack)

	log.Printf("%s %v %s \n ** Call Trace ** \n%s\n", a...)
}

func Fatal(msg string, arg ...any) {
	caller := rt.CallerInfo(runtime.Caller(1))
	a := make([]any, 0, 3+len(arg))
	a = append(a, ts(), red.Sprint("FATAL"), caller)
	a = append(a, arg...)

	log.Printf("%s %v %s "+msg, a...)

	dump(caller, fmt.Sprintf(msg, arg...), "")
}
