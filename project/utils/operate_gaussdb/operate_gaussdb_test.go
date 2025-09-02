package operategaussdb

import (
	"gin_template/project/utils/logger"
	"testing"
)

func TestConn(t *testing.T) {
	logger.InitLogger(false)
	InitGaussDB(nil)
	Close()
}
