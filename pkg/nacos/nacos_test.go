package nacos

import (
	"fmt"
	"gin_template/pkg/deployenv"
	"testing"
)

func TestNacosLogin(t *testing.T) {
	env := InitNacos(deployenv.DEV)
	fmt.Println(env)
	token, err := NacosCli.Login()
	fmt.Println(token, err)
	cfg, ty, err := NacosCli.GetConfig(token, "mongodb-standalone-test-test")
	fmt.Println(cfg, ty, err)
}
