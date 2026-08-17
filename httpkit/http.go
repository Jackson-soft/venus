package httpkit

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"time"
)

// 定制化的http客户端

const (
	dialTimeout           = 90 * time.Second // 拨号超时
	tlsHandshakeTimeout   = 90 * time.Second // TLS握手超时
	responseHeaderTimeout = 90 * time.Second // 响应头超时
	maxIdleConnsPerHost   = 16               // 同一主机的最大空闲连接数
	requestTimeout        = 90 * time.Second // 全局请求超时
)

var DefaultClient = &http.Client{
	Transport: &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: dialTimeout,
		}).DialContext,
		TLSHandshakeTimeout:   tlsHandshakeTimeout,
		ResponseHeaderTimeout: responseHeaderTimeout,
		MaxIdleConnsPerHost:   maxIdleConnsPerHost,
	},
	Timeout: requestTimeout,
}

type WebBase struct {
	Context context.Context   // 请求上下文，用于取消请求或传递链路信息
	Url     string            // 请求地址
	Method  string            // 请求方法
	Header  map[string]string // 请求头
	Body    io.Reader         // 请求体
}

// http 请求的工具函数
func WebDo(base *WebBase) ([]byte, error) {
	ctx := base.Context
	if ctx == nil {
		ctx = context.Background()
	}

	req, err := http.NewRequestWithContext(ctx, base.Method, base.Url, base.Body)
	if err != nil {
		return nil, err
	}

	for key, value := range base.Header {
		req.Header.Set(key, value)
	}

	resp, err := DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, errors.New(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func HttpDo(base *WebBase, resp any) error {
	body, err := WebDo(base)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, resp)
}
