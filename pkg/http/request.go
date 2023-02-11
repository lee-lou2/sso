package http

import (
	"io"
	"io/ioutil"
	netHttp "net/http"
)

// Response 요청 결과 값
type Response struct {
	StatusCode int
	Body       string
}

// Header 헤더 데이터
type Header struct {
	Key   string
	Value string
}

// Request 요청
func Request(method string, url string, payload io.Reader, headers ...*Header) (*Response, error) {
	client := &netHttp.Client{}

	req, err := netHttp.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	// 헤더 값 지정
	if len(headers) > 0 {
		for _, header := range headers {
			req.Header.Set(header.Key, header.Value)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		StatusCode: resp.StatusCode,
		Body:       string(body),
	}, nil
}
