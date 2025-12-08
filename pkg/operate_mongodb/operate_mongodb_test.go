package operatemongodb

import (
	"gin_template/pkg/logger"
	"testing"
)

func TestConn(t *testing.T) {
	logger.InitLogger(false)
	InitMongoDB(nil)
	Close()
}
