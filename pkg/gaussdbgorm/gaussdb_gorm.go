package gaussdbgorm

import (
	"context"
	"database/sql"
	"gin_template/pkg/logger"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/things-go/gormzap"
	"gorm.io/driver/gaussdb"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	glog "gorm.io/gorm/logger"
)

var (
	GDB       *gorm.DB
	DB        *sql.DB
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

func guassDBDNSBuilder(g *GaussDBConf) string {
	dsnBuilder := strings.Builder{}
	dsnBuilder.WriteString("host=")
	dsnBuilder.WriteString(g.Host)
	dsnBuilder.WriteString(" port=")
	dsnBuilder.WriteString(strconv.Itoa(g.Port))
	dsnBuilder.WriteString(" user=")
	dsnBuilder.WriteString(g.User)
	dsnBuilder.WriteString(" password=")
	dsnBuilder.WriteString(g.Password)
	dsnBuilder.WriteString(" dbname=")
	dsnBuilder.WriteString(g.DbName)
	dsnBuilder.WriteString(" sslmode=")
	dsnBuilder.WriteString(g.Sslmode)
	dsnBuilder.WriteString(" TimeZone=")
	dsnBuilder.WriteString("Asia/Shanghai")
	return dsnBuilder.String()
}

// 连接 GaussDB
func InitGaussDBGorm(g *GaussDBConf) error {
	var err error
	onceConn.Do(func() {
		err = connectGaussDB(g)
	})
	return err
}

func connectGaussDB(g *GaussDBConf) error {
	if g == nil {
		g = &GaussDBConf{
			Host:     "127.0.0.1",
			Port:     5432,
			User:     "gaussdb",
			Password: "Gaussdb@1",
			DbName:   "postgres",
			Sslmode:  "disable",
		}
	}
	var err error
	l := gormzap.New(logger.GetLogger())
	GDB, err = gorm.Open(gaussdb.New(gaussdb.Config{
		DSN: guassDBDNSBuilder(g),
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用复数表名
		},
		Logger: l.LogMode(glog.Warn),
	})
	if err != nil {
		logger.Errorf("GaussDB connect error: %s", err.Error())
		return err
	}
	DB, err = GDB.DB()
	if err != nil {
		logger.Errorf("GaussDB get DB error: %s", err.Error())
		return err
	}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()
	err = DB.PingContext(ctx)
	if err != nil {
		logger.Errorf("GaussDB ping error: %s", err.Error())
		return err
	}
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	DB.SetMaxIdleConns(3)
	// SetMaxOpenCons 设置数据库的最大连接数量。
	DB.SetMaxOpenConns(10)
	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	DB.SetConnMaxLifetime(time.Hour)
	logger.Info("GaussDB connect")
	return nil
}

// 关闭 GaussDB
func CloseGaussDB() error {
	var err error
	onceClose.Do(func() {
		if DB != nil {
			err = DB.Close()
		}
	})
	if err != nil {
		logger.Warnf("GaussDB close error: %s", err.Error())
		return err
	} else {
		logger.Info("GaussDB closed")
		return nil
	}
}
