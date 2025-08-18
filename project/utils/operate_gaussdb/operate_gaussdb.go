package operategaussdb

import (
	"context"
	"database/sql"
	"fmt"
	"gin_template/project/config"
	"gin_template/project/utils/logger"
	"time"

	_ "gitee.com/opengauss/openGauss-connector-go-pq"
)

var GaussDB *sql.DB

// 根据配置初始化连接mongodb
func InitGaussDB() (err error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Cfg.GaussDB.Host, config.Cfg.GaussDB.Port, config.Cfg.GaussDB.User,
		config.Cfg.GaussDB.Password, config.Cfg.GaussDB.DbName, config.Cfg.GaussDB.Sslmode)
	GaussDB, err = sql.Open("opengauss", connStr)
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()
	err = GaussDB.PingContext(ctx)
	if err != nil {
		return
	}
	logger.Info("GaussDB connect...")
	return
}

func Close() (err error) {
	err = GaussDB.Close()
	if err != nil {
		return
	}
	logger.Infof("GaussDB closed")
	return
}
