package db

import (
	"testing"
)

func TestInit(t *testing.T) {
	err := Init("/tmp/stargazer.db")
	if err != nil {
		t.Errorf("failed to init database: %v", err)
	}
	return
}
