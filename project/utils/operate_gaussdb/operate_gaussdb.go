package operategaussdb

import (
	"context"
	"database/sql"
	"fmt"
	"gin_template/project/utils/logger"
	"time"

	_ "gitee.com/opengauss/openGauss-connector-go-pq"
)

var GaussDB *sql.DB

type GaussDBConf struct {
	Host     string `yaml:"HOST"`
	Port     int    `yaml:"PORT"`
	User     string `yaml:"USER"`
	Password string `yaml:"PASSWORD"`
	DbName   string `yaml:"DATABASE"`
	Sslmode  string `yaml:"SSLMODE"`
}

// 根据配置初始化连接mongodb
func InitGaussDB(gsCfg *GaussDBConf) (err error) {
	if gsCfg == nil {
		gsCfg = &GaussDBConf{
			Host:     "127.0.0.1",
			Port:     5432,
			User:     "test",
			Password: "Gaussdb@1",
			DbName:   "postgres",
			Sslmode:  "disable",
		}
	}
	if gsCfg.Sslmode == "" {
		gsCfg.Sslmode = "disable"
	}
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		gsCfg.Host, gsCfg.Port, gsCfg.User,
		gsCfg.Password, gsCfg.DbName, gsCfg.Sslmode)
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
