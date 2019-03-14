package logger

import (
	"fmt"
	"time"
)

func LogIfErr(err error, a ...interface{}) {
	if err != nil {
		Error(err, a)
	}
}

func Error(err error, a ...interface{}) {
	log("ERROR", err.Error(), a...)
}

func Warn(err error, a ...interface{}) {
	log("WARN", err.Error(), a...)
}

func Info(format interface{}, a ...interface{}) {
	log("INFO", interfaceToString(format), a...)
}

func log(logType string, format string, a ...interface{}) {
	now := time.Now()
	fmt.Println("[" + logType + "] " + now.Format("2006-01-02 15:04:05") + " | " + fmt.Sprintf(format, a...))
}

func interfaceToString(message interface{}) string {
	if v, ok := message.(error); ok {
		return v.Error()
	}
	return fmt.Sprint(message)
}
