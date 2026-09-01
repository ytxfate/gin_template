package gaussdb

import (
	"gin_template/pkg/logger"
	"testing"
)

func TestConn(t *testing.T) {
	logger.InitLogger(false)
	err := InitGaussDB(nil)
	if err != nil {
		t.Error(err)
		return
	}
	err = Close()
	if err != nil {
		t.Error(err)
		return
	}
}
