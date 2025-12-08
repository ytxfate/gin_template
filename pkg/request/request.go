package request

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"time"
)

// 请求
//
// 注: json 优先级高于 body
type Request struct {
	Ctx     context.Context
	Url     string
	Method  string
	Params  [][2]string
	Body    io.Reader
	Json    interface{}
	Headers map[string][]string
	Timeout time.Duration
}

type option func(*Request)

// 初始化请求
func NewRequest(url string, opts ...option) *Request {
	req := &Request{
		Ctx:     context.Background(),
		Url:     url,
		Method:  http.MethodGet,
		Params:  nil,
		Body:    nil,
		Json:    nil,
		Headers: nil,
		Timeout: time.Second * 60,
	}
	for _, opt := range opts {
		opt(req)
	}
	return req
}

func WithCtx(ctx context.Context) option {
	return func(r *Request) {
		r.Ctx = ctx
	}
}

func WithMethod(method string) option {
	return func(r *Request) {
		r.Method = method
	}
}

func WithParams(params [][2]string) option {
	return func(r *Request) {
		r.Params = params
	}
}

func WithBody(body io.Reader) option {
	return func(r *Request) {
		r.Body = body
	}
}

func WithJson(j interface{}) option {
	return func(r *Request) {
		r.Json = j
	}
}

func WithHeaders(headers map[string][]string) option {
	return func(r *Request) {
		r.Headers = headers
	}
}

func WithTimeout(timeout time.Duration) option {
	return func(r *Request) {
		r.Timeout = timeout
	}
}

func (request *Request) Do() (resp *http.Response, err error) {
	if request.Method == "" {
		request.Method = http.MethodGet
	}
	// 请求接口获取token
	var req *http.Request
	body := request.Body
	if request.Json != nil {
		by := new(bytes.Buffer)
		err = json.NewEncoder(by).Encode(request.Json)
		if err != nil {
			return
		}
		body = bufio.NewReader(by)
	}
	req, err = http.NewRequestWithContext(request.Ctx, request.Method, request.Url, body)
	if err != nil {
		return
	}
	// headers
	for hKey, hVals := range request.Headers {
		for _, hVal := range hVals {
			req.Header.Add(hKey, hVal)
		}
	}
	if request.Json != nil {
		req.Header.Del("content-type")
		req.Header.Add("content-type", "application/json")
	}
	// params
	qry := req.URL.Query()
	for _, param := range request.Params {
		qry.Add(param[0], param[1])
	}
	req.URL.RawQuery = qry.Encode()

	client := http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: request.Timeout, // TCP连接超时
			}).DialContext,
			ResponseHeaderTimeout: request.Timeout, // 等待响应头
			IdleConnTimeout:       request.Timeout, // 空闲连接回收
		},
	}
	resp, err = client.Do(req)
	if err != nil {
		return
	}
	return
}
