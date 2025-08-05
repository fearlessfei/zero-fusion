package httpclient

import (
	"io"
	"net/http"
	"strings"
)

// Response 响应结构
type Response struct {
	RawResponse *http.Response

	body []byte
	size int64
}

// SetBody 设置body
func (r *Response) SetBody(body []byte) {
	r.body = body
}

// Body 响应返回为已执行请求的字节数组
func (r *Response) Body() []byte {
	if r.RawResponse == nil {
		return []byte{}
	}
	return r.body
}

// Status 返回已执行请求的HTTP状态字符串
//
//	Example: 200 OK
func (r *Response) Status() string {
	if r.RawResponse == nil {
		return ""
	}
	return r.RawResponse.Status
}

// StatusCode 返回已执行请求的HTTP状态代码
//
//	Example: 200
func (r *Response) StatusCode() int {
	if r.RawResponse == nil {
		return 0
	}
	return r.RawResponse.StatusCode
}

// Header 响应头
func (r *Response) Header() http.Header {
	if r.RawResponse == nil {
		return http.Header{}
	}
	return r.RawResponse.Header
}

// Cookies 响应cookies
func (r *Response) Cookies() []*http.Cookie {
	if r.RawResponse == nil {
		return make([]*http.Cookie, 0)
	}
	return r.RawResponse.Cookies()
}

// String body作为字符串返回
func (r *Response) String() string {
	if r.body == nil {
		return ""
	}
	return strings.TrimSpace(string(r.body))
}

// SetSize 设置size
func (r *Response) SetSize(size int64) {
	r.size = size
}

// Size 以字节为单位返回HTTP响应大小
func (r *Response) Size() int64 {
	return r.size
}

// RawBody 原始响应body
func (r *Response) RawBody() io.ReadCloser {
	if r.RawResponse == nil {
		return nil
	}
	return r.RawResponse.Body
}

// IsSuccess 如果状态码 `code >= 200 and <= 299` 返回true其他返回false
func (r *Response) IsSuccess() bool {
	return r.StatusCode() > 199 && r.StatusCode() < 300
}

// IsError 状态码 `code >= 400`返回true其他返回false
func (r *Response) IsError() bool {
	return r.StatusCode() > 399
}
