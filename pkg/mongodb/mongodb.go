package mongodb

import (
	"context"
	"gin_template/pkg/logger"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var (
	MgClient  *mongo.Client
	MgDB      *mongo.Database
	onceConn  sync.Once
	onceClose sync.Once
)

type MongodbConf struct {
	Url       string `yaml:"URL"`        // 有此项则优先用此项进行数据库连接否则用 HOST 和 PORT 连接
	Username  string `yaml:"USERNAME"`   // 用户名
	Password  string `yaml:"PASSWORD"`   // 密码
	DefaultDb string `yaml:"DEFAULT_DB"` // 默认数据库
}

// 根据配置初始化连接mongodb
func InitMongoDB(mgCfg *MongodbConf) error {
	var err error
	onceConn.Do(func() {
		err = initMongoDB(mgCfg)
	})
	return err
}

func initMongoDB(mgCfg *MongodbConf) error {
	var err error
	if mgCfg == nil {
		mgCfg = &MongodbConf{
			Url:       "127.0.0.1:27017",
			Username:  "test",
			Password:  "test",
			DefaultDb: "test",
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	opts := options.ClientOptions{}
	opts.SetAuth(options.Credential{
		AuthSource: mgCfg.DefaultDb,
		Username:   mgCfg.Username,
		Password:   mgCfg.Password,
	})
	opts.SetHosts(strings.Split(mgCfg.Url, ","))
	opts.SetMaxPoolSize(10) // 设置最大连接池
	MgClient, err = mongo.Connect(&opts)
	if err != nil {
		logger.Errorf("MongoDB connect err: %v", err)
		return err
	}
	err = MgClient.Ping(ctx, readpref.Primary())
	if err != nil {
		logger.Errorf("MongoDB ping err: %v", err)
		return err
	}
	MgDB = MgClient.Database(mgCfg.DefaultDb)
	logger.Info("Mongodb Connect..")
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
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()
	err := MgClient.Disconnect(ctx)
	if err != nil {
		logger.Errorf("MongoDB ping err: %v", err)
	}
	logger.Info("Mongodb closed")
	return nil
}
