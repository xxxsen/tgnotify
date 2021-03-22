package log

import (
	"fmt"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

//refer:https://gist.github.com/miguelmota/0508139b2df9142b4574dba2edb6fb9b

//FileLineHook FileLineHook
type FileLineHook struct{}

//Levels levels
func (hook FileLineHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire ...
func (hook FileLineHook) Fire(entry *logrus.Entry) error {
	if _, file, line, ok := runtime.Caller(9); ok {
		entry.Data["loc"] = fmt.Sprintf("%s:%v", path.Base(file), line)
	}
	return nil
}
