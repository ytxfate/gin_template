package gaussdb

import (
	"context"
	"database/sql"
	"fmt"
	"gin_template/pkg/logger"
	"sync"
	"time"

	_ "gitee.com/opengauss/openGauss-connector-go-pq"
)

var (
	GaussDB   *sql.DB
	onceConn  sync.Once
	onceClose sync.Once
)

type GaussDBConf struct {
	Host     string `yaml:"HOST"`
	Port     int    `yaml:"PORT"`
	User     string `yaml:"USER"`
	Password string `yaml:"PASSWORD"`
	DbName   string `yaml:"DATABASE"`
	Sslmode  string `yaml:"SSLMODE"`
}

// 根据配置初始化连接mongodb
func InitGaussDB(gsCfg *GaussDBConf) error {
	var err error
	onceConn.Do(func() {
		err = initGaussDB(gsCfg)
	})
	return err
}

func initGaussDB(gsCfg *GaussDBConf) error {
	var err error
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
		return err
	}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()
	err = GaussDB.PingContext(ctx)
	if err != nil {
		return err
	}
	logger.Info("GaussDB connect...")
	return err
}

func Close() error {
	var err error
	onceClose.Do(func() {
		err = close()
	})
	return err
}

func close() error {
	err := GaussDB.Close()
	if err != nil {
		return err
	}
	logger.Infof("GaussDB closed")
	return nil
}
