package configs

import (
	"gin_template/pkg/deployenv"
	"gin_template/pkg/nacos"
	"testing"
)

func TestDBConfig(t *testing.T) {
	env := nacos.InitNacos(deployenv.DEV)
	err := InitAllDBConfig(env)
	if err != nil {
		t.Error(err)
		return
	}
}
