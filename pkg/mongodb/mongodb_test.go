package mongodb

import (
	"gin_template/pkg/logger"
	"testing"
)

func TestConn(t *testing.T) {
	logger.InitLogger(false)
	err := InitMongoDB(nil)
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
