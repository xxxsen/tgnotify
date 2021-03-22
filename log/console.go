package log

import "fmt"

//ConsoleWriter 控制台写类
type ConsoleWriter struct {
}

//Write 写数据到控制台
func (cw *ConsoleWriter) Write(p []byte) (n int, err error) {
	return fmt.Print(string(p))
}
