package logger

import (
	"fmt"
	"log"
	"runtime/debug"
)

func LogIfErr(err error, a ...interface{}) {
	if err != nil {
		Error(err, a)
	}
}

func Error(err error, a ...interface{}) {
	logs("ERROR", err.Error(), a...)
}

func Warn(err error, a ...interface{}) {
	logs("WARN", err.Error(), a...)
}

func Info(format interface{}, a ...interface{}) {
	logs("INFO", interfaceToString(format), a...)
}

func logs(logType string, format string, a ...interface{}) {
	if logType == "ERROR" {
		format += "\nStacktrace:\n" + string(debug.Stack())
	}

	log.Printf("[" + logType + "] "+ fmt.Sprintf(format, a...))
}

func interfaceToString(message interface{}) string {
	if v, ok := message.(error); ok {
		return v.Error()
	}
	return fmt.Sprint(message)
}
