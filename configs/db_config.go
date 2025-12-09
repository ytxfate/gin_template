package configs

import (
	"gin_template/pkg/deployenv"
	gaussdb "gin_template/pkg/gaussdb"
	mongodb "gin_template/pkg/mongodb"
	"gin_template/pkg/nacos"

	"gopkg.in/yaml.v3"
)

var MgConf *mongodb.MongodbConf

func initMgConf(env deployenv.DeployEnv, nacosToken string) {
	dataId := "mongodb-standalone-test-test"
	if env == deployenv.PROD {
		dataId = "mongodb-rs-database-user"
	}
	cfg, _, err := nacos.NacosCfg.GetConfig(nacosToken, dataId)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(cfg), &MgConf)
	if err != nil {
		panic(err)
	}
}

var GaussCfg *gaussdb.GaussDBConf

func initGaussCfg(env deployenv.DeployEnv, nacosToken string) {
	dataId := "gaussdb-standalone-test-test"
	if env == deployenv.PROD {
		dataId = "gaussdb-rs-database-user"
	}
	cfg, _, err := nacos.NacosCfg.GetConfig(nacosToken, dataId)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(cfg), &GaussCfg)
	if err != nil {
		panic(err)
	}
}

func InitAllDBConfig(env deployenv.DeployEnv) error {
	token, err := nacos.NacosCfg.Login()
	if err != nil {
		return err
	}
	initMgConf(env, token)
	initGaussCfg(env, token)
	return nil
}
