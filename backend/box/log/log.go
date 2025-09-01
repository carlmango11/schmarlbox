package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

var Debug = false
var out *os.File

func init() {
	f, err := os.OpenFile("output.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	out = f
}

func Println(s string) {
	logf(s)
}

func Printf(format string, v ...interface{}) {
	logf(format, v...)
}

func Panicf(format string, v ...interface{}) {
	logf(format, v...)
	log.Panicf(format, v...)
}

func Debugf(format string, v ...interface{}) {
	if Debug {
		logf(format, v...)
	}
}

func logf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...) + "\n"
	s = fmt.Sprintf("%v %v", time.Now().Format(time.RFC3339), s)

	out.WriteString(s)
	//log.Printf(format, v...)
}
