package database

import (
	"github.com/sptuan/stargazer/modules/logger"
	"testing"
)

func TestInit(t *testing.T) {
	logger.Init()
	err := Init("/tmp/stargazer.db")
	if err != nil {
		t.Errorf("failed to init database: %v", err)
	}
	return
}
