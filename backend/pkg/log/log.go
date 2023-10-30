package log

import (
	"github.com/golang/glog"
)

func Info(template string, args ...interface{}) {
	glog.Info(template, args)
}
func Warning(template string, args ...interface{}) {
	glog.Warning(template, args)
}
func Error(template string, args ...interface{}) {
	glog.Error(template, args)
}
func Fatal(template string, args ...interface{}) {
	glog.Fatal(template, args)
}
