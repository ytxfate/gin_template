package config

import (
	gaussdb "gin_template/pkg/gaussdb"
	mongodb "gin_template/pkg/mongodb"

	"gopkg.in/yaml.v3"
)

var MgConf *mongodb.MongodbConf

func initMgConf(env DeployEnv, nacosCfg NacosServerConfig, nacosToken string) {
	dataId := "mongodb-standalone-test-test"
	if env == PROD {
		dataId = "mongodb-rs-database-user"
	}
	cfg, _, err := nacosCfg.getNacosConfig(nacosToken, dataId)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(cfg), &MgConf)
	if err != nil {
		panic(err)
	}
}

var GaussCfg *gaussdb.GaussDBConf

func initGaussCfg(env DeployEnv, nacosCfg NacosServerConfig, nacosToken string) {
	dataId := "gaussdb-standalone-test-test"
	if env == PROD {
		dataId = "gaussdb-rs-database-user"
	}
	cfg, _, err := nacosCfg.getNacosConfig(nacosToken, dataId)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(cfg), &GaussCfg)
	if err != nil {
		panic(err)
	}
}

func InitAllDBConfig(env DeployEnv, nacosCfg NacosServerConfig, nacosToken string) {
	initMgConf(env, nacosCfg, nacosToken)
	initGaussCfg(env, nacosCfg, nacosToken)
}
