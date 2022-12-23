package config

import "testing"

func TestConfigInit(t *testing.T) {
	Init("")
}

func TestConfigInitWithPath(t *testing.T) {
	Init("../../config")
}
