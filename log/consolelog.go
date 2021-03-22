package log

import (
	"fmt"
	"os"
)

//ConsoleWrite 写日志到console
func ConsoleWrite(formatter string, args ...interface{}) {
	fmt.Printf(formatter+"\n", args...)
}

//ConsoleError console错误日志
func ConsoleError(formatter string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, formatter+"\n", args...)
}
