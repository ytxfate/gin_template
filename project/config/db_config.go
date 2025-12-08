package config

import (
	gaussdb "gin_template/pkg/gaussdb"
	mongodb "gin_template/pkg/mongodb"

	"gopkg.in/yaml.v3"
)

var MgConf *mongodb.MongodbConf

func initMgConf() {
	dataId := "mongodb-standalone-test-test"
	if Cfg.Env == PROD {
		dataId = "mongodb-rs-database-user"
	}
	cfg, _, err := Cfg.nacosCfg.getNacosConfig(Cfg.nacosToken, dataId)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(cfg), &MgConf)
	if err != nil {
		panic(err)
	}
}

var GaussCfg *gaussdb.GaussDBConf

func initGaussCfg() {
	dataId := "gaussdb-standalone-test-test"
	if Cfg.Env == PROD {
		dataId = "gaussdb-rs-database-user"
	}
	cfg, _, err := Cfg.nacosCfg.getNacosConfig(Cfg.nacosToken, dataId)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(cfg), &GaussCfg)
	if err != nil {
		panic(err)
	}
}

func initAllConfig() {
	initMgConf()
	initGaussCfg()
}
