package obmysqlgorm

import (
	"context"
	"database/sql"
	"gin_template/pkg/logger"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/things-go/gormzap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glog "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	GDB       *gorm.DB
	DB        *sql.DB
	onceConn  sync.Once
	onceClose sync.Once
)

type OceanBaseConf struct {
	Host     string `yaml:"HOST"`
	Port     int    `yaml:"PORT"`
	User     string `yaml:"USER"`
	Password string `yaml:"PASSWORD"`
	DbName   string `yaml:"DATABASE"`
}

func oceanBaseDNSBuilder(o *OceanBaseConf) string {
	// dsn := "user_name:******@tcp(host:port)/schema_name?charset=utf8mb4&parseTime=True&loc=Local"
	dsnBuilder := strings.Builder{}
	dsnBuilder.WriteString(o.User)
	dsnBuilder.WriteString(":")
	dsnBuilder.WriteString(o.Password)
	dsnBuilder.WriteString("@tcp(")
	dsnBuilder.WriteString(o.Host)
	dsnBuilder.WriteString(":")
	dsnBuilder.WriteString(strconv.Itoa(o.Port))
	dsnBuilder.WriteString(")/")
	dsnBuilder.WriteString(o.DbName)
	dsnBuilder.WriteString("?charset=utf8mb4&parseTime=True&loc=Local")
	return dsnBuilder.String()
}

// 连接 OceanBase
func InitOceanBaseGorm(g *OceanBaseConf) error {
	var err error
	onceConn.Do(func() {
		err = connectOceanBase(g)
	})
	return err
}

func connectOceanBase(o *OceanBaseConf) error {
	if o == nil {
		o = &OceanBaseConf{
			Host:     "127.0.0.1",
			Port:     5432,
			User:     "OceanBase",
			Password: "OceanBase@1",
			DbName:   "postgres",
		}
	}
	var err error
	l := gormzap.New(logger.GetLogger(), gormzap.WithConfig(glog.Config{
		SlowThreshold: time.Second, // 一秒以上才算慢SQL
		LogLevel:      glog.Warn,   // 日志等级
	}))
	GDB, err = gorm.Open(mysql.Open(oceanBaseDNSBuilder(o)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用复数表名
		},
		Logger: l,
	})
	if err != nil {
		logger.Errorf("OceanBase connect error: %s", err.Error())
		return err
	}
	DB, err = GDB.DB()
	if err != nil {
		logger.Errorf("OceanBase get DB error: %s", err.Error())
		return err
	}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()
	err = DB.PingContext(ctx)
	if err != nil {
		logger.Errorf("OceanBase ping error: %s", err.Error())
		return err
	}
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	DB.SetMaxIdleConns(3)
	// SetMaxOpenCons 设置数据库的最大连接数量。
	DB.SetMaxOpenConns(10)
	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	DB.SetConnMaxLifetime(time.Hour)
	logger.Info("OceanBase connect")
	return nil
}

// 关闭 OceanBase
func CloseOceanBase() error {
	var err error
	onceClose.Do(func() {
		if DB != nil {
			err = DB.Close()
		}
	})
	if err != nil {
		logger.Warnf("OceanBase close error: %s", err.Error())
		return err
	} else {
		logger.Info("OceanBase closed")
		return nil
	}
}
