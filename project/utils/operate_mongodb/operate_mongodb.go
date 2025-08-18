package operatemongodb

import (
	"context"
	"gin_template/project/config"
	"gin_template/project/utils/logger"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var MgClient *mongo.Client
var MgDB *mongo.Database

// 根据配置初始化连接mongodb
func InitMongoDB() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	opts := options.ClientOptions{}
	opts.SetAuth(options.Credential{
		AuthSource: config.Cfg.MongoDB.DefaultDb,
		Username:   config.Cfg.MongoDB.Username,
		Password:   config.Cfg.MongoDB.Password,
	})
	opts.SetHosts(strings.Split(config.Cfg.MongoDB.Url, ","))
	opts.SetMaxPoolSize(10) // 设置最大连接池
	MgClient, err = mongo.Connect(&opts)
	if err != nil {
		logger.Errorf("MongoDB connect err: %v", err)
		return
	}
	err = MgClient.Ping(ctx, readpref.Primary())
	if err != nil {
		logger.Errorf("MongoDB ping err: %v", err)
		return
	}
	MgDB = MgClient.Database(config.Cfg.MongoDB.DefaultDb)
	logger.Info("Mongodb Connect..")
	return
}

func Close() (err error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()
	err = MgClient.Disconnect(ctx)
	if err != nil {
		logger.Errorf("MongoDB ping err: %v", err)
	}
	logger.Info("Mongodb closed")
	return
}
