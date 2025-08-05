package httpclient

import (
	"context"
)

// HTTPClient http客户端
type HTTPClient interface {
	// Get get请求
	Get(ctx context.Context, url string, body any) (*Response, error)
	// Post post请求
	Post(ctx context.Context, url string, body any) (*Response, error)
	// Put put请求
	Put(ctx context.Context, url string, body any) (*Response, error)
	// Delete delete请求
	Delete(ctx context.Context, url string, body any) (*Response, error)
	// Do 做请求
	Do(ctx context.Context, method string, url string, body any) (*Response, error)
}
