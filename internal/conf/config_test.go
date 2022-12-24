package conf

import "testing"

func TestConfigInit(t *testing.T) {
	Init("")
}

func TestConfigInitWithPath(t *testing.T) {
	Init("../../config/config.yaml")
	t.Logf("Config: %+v", Cfg)
}
