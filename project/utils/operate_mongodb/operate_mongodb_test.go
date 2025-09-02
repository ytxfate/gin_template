package operatemongodb

import (
	"gin_template/project/utils/logger"
	"testing"
)

func TestConn(t *testing.T) {
	logger.InitLogger(false)
	InitMongoDB(nil)
	Close()
}
