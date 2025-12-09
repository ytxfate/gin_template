package request

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"
)

var pubHeaders = map[string][]string{
	"User-Agent": {"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36"},
}

func TestGet(t *testing.T) {
	req := NewRequest("https://jsonplaceholder.org/users",
		WithParams([][2]string{{"id", "1"}}),
		WithHeaders(pubHeaders),
		WithCtx(context.Background()),
		WithMethod(""),
		WithTimeout(time.Second*5),
	)
	resp, err := req.Do()
	if err != nil {
		fmt.Println("req err:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("http status: ", resp.StatusCode, resp.Status)
	var data map[string]interface{}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("resp read err:", err)
		return
	}
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		fmt.Println("resp json decode err:", err, string(respBody))
		return
	}
	fmt.Println(data)
}

func TestPost(t *testing.T) {
	req := NewRequest("https://open.feishu.cn/open-apis/passport/v1/sessions/query",
		WithMethod(http.MethodPost),
		WithHeaders(pubHeaders),
		WithBody(bytes.NewBufferString("{\"a\": 1}")),
		WithJson(map[string][]string{
			"user_ids": {
				"7aeddb6a1",
				"7aeddb6a2",
				"7aeddb6a3",
			},
		}),
	)
	resp, err := req.Do()
	if err != nil {
		fmt.Println("req err:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(req)
	fmt.Println("http status: ", resp.StatusCode, resp.Status)
	var data map[string]interface{}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("resp read err:", err)
		return
	}
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		fmt.Println("resp json decode err:", err, string(respBody))
		return
	}
	fmt.Println(data)
}
