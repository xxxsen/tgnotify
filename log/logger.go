package log

import (
	"io"
	"strings"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logInst *logrus.Logger

func init() {
	logInst = logrus.New()
}

const (
	//LogPanic panic
	LogPanic = iota
	//LogFatal fatal
	LogFatal
	//LogError error
	LogError
	//LogWarn warn
	LogWarn
	//LogInfo info
	LogInfo
	//LogDebug debug
	LogDebug
	//LogTrace trace
	LogTrace
)

//Instance 获取全局实例
func Instance() *logrus.Logger {
	return logInst
}

//StringToLevel 日志级别转换
func StringToLevel(lv string) int {
	lolv := strings.ToLower(lv)
	switch lolv {
	case "panic":
		return LogPanic
	case "fatal":
		return LogFatal
	case "warn":
		return LogWarn
	case "info":
		return LogInfo
	case "debug":
		return LogDebug
	case "trace":
		return LogTrace
	}
	return LogTrace
}

func getLevel(lv int) logrus.Level {
	return logrus.Level(lv)
}

type nullWriter struct {
}

func (nw *nullWriter) Write(data []byte) (int, error) {
	return len(data), nil
}

//UTCFormatter 输出utc时间
type UTCFormatter struct {
	logrus.Formatter
}

//Format Format
func (u UTCFormatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = e.Time.Local()
	return u.Formatter.Format(e)
}

//Init 日志初始化
func Init(file string, lv int, maxRotate int, maxSizeMB int, maxKeepDays int, withConsole bool) {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05.000"
	customFormatter.FullTimestamp = true
	logInst.SetFormatter(&UTCFormatter{customFormatter})
	var logger io.Writer = &nullWriter{}
	if len(file) != 0 && maxSizeMB > 0 {
		logger = &lumberjack.Logger{
			// 日志输出文件路径
			Filename:   file,
			MaxSize:    maxSizeMB, // megabytes
			MaxBackups: maxRotate,
			MaxAge:     maxKeepDays, //days
			Compress:   false,       // disabled by default
		}
	}
	if withConsole {
		logger = io.MultiWriter(&ConsoleWriter{}, logger)
	}
	logInst.SetOutput(logger)
	logInst.AddHook(FileLineHook{})
	logInst.SetLevel(getLevel(lv))
}

//Fatal Fatal
func Fatal(args ...interface{}) {
	logInst.Fatal(args...)
}

//Error Error
func Error(args ...interface{}) {
	logInst.Error(args...)
}

//Info Info
func Info(args ...interface{}) {
	logInst.Info(args...)
}

//Debug Debug
func Debug(args ...interface{}) {
	logInst.Debug(args...)
}

//Trace Trace
func Trace(args ...interface{}) {
	logInst.Trace(args...)
}

//Fatalf Fatalf
func Fatalf(formatter string, args ...interface{}) {
	logInst.Fatalf(formatter, args...)
}

//Errorf Errorf
func Errorf(formatter string, args ...interface{}) {
	logInst.Errorf(formatter, args...)
}

//Infof Infof
func Infof(formatter string, args ...interface{}) {
	logInst.Infof(formatter, args...)
}

//Debugf Debugf
func Debugf(formatter string, args ...interface{}) {
	logInst.Debugf(formatter, args...)
}

//Tracef Tracef
func Tracef(formatter string, args ...interface{}) {
	logInst.Tracef(formatter, args...)
}
