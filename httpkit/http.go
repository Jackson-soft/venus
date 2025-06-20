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

var DefaultClient = &http.Client{
	Transport: &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 90 * time.Second, // 拨号超时
		}).DialContext,
		TLSHandshakeTimeout:   90 * time.Second, // TLS握手超时
		ResponseHeaderTimeout: 90 * time.Second, // 响应头超时
		MaxIdleConnsPerHost:   16,               // 同一主机的最大空闲连接数
	},
	Timeout: 90 * time.Second, // 全局请求超时
}

type WebBase struct {
	Url    string            // 请求地址
	Method string            // 请求方法
	Header map[string]string // 请求头
	Body   io.Reader         // 请求体
}

// http 请求的工具函数
func WebDo(base *WebBase) ([]byte, error) {
	req, err := http.NewRequestWithContext(context.Background(), base.Method, base.Url, base.Body)
	if err != nil {
		return nil, err
	}

	for key, value := range base.Header {
		req.Header.Add(key, value)
	}

	resp, err := DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
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
