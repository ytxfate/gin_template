package config

import (
	"fmt"
	"testing"
)

func TestNacosLogin(t *testing.T) {
	nacosCfg := NewNacosServerConfigTest()
	token, err := nacosCfg.NacosLogin()
	fmt.Println(token, err)
}
