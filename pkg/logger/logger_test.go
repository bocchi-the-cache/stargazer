package logger

import "testing"

func TestLogger(t *testing.T) {
	Init()
	Infof("hello")
	Warnf("warn %s", "world")
	Errorf("error %s", "world")
	Panicf("panic %s", "world")
}
