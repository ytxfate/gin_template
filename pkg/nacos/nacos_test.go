package nacos

import (
	"fmt"
	"gin_template/pkg/deployenv"
	"testing"
)

func TestNacosLogin(t *testing.T) {
	env := InitNacos(deployenv.DEV)
	fmt.Println(env)
	token, err := NacosCfg.Login()
	fmt.Println(token, err)
	cfg, ty, err := NacosCfg.GetConfig(token, "mongodb-standalone-test-test")
	fmt.Println(cfg, ty, err)
}
