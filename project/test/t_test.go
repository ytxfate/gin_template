package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type tt struct {
	ServerIP string `json:"server_Ip"`
}

func TestStructTags(t *testing.T) {
	var s1 string = "{\"SERVERIP\": \"t\"}"
	var s2 string = "{\"SERVER_IP\": \"T\"}"
	var tt1 tt
	var tt2 tt
	_ = json.Unmarshal([]byte(s1), &tt1)
	fmt.Printf("%#v\n", tt1)
	_ = json.Unmarshal([]byte(s2), &tt2)
	fmt.Printf("%#v\n", tt2)
}
