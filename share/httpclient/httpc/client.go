package httpc

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest/httpc"

	"zero-fusion/share/httpclient"
)

const defaultTimeout = 3 * time.Second

type client struct {
	name      string
	RawClient *http.Client
	service   httpc.Service
}

func NewHTTPClient(name string, opts ...httpc.Option) httpclient.HTTPClient {
	rawClient := &http.Client{
		Timeout: defaultTimeout,
	}
	service := httpc.NewServiceWithClient(name, rawClient, opts...)

	return &client{
		name:      name,
		RawClient: rawClient,
		service:   service,
	}
}

func NewHTTPClientWithClient(name string, cli *http.Client, opts ...httpc.Option) httpclient.HTTPClient {
	service := httpc.NewServiceWithClient(name, cli, opts...)

	return &client{
		RawClient: cli,
		service:   service,
	}
}

// Get get请求
func (c *client) Get(ctx context.Context, url string, body any) (*httpclient.Response, error) {
	return c.Do(ctx, http.MethodGet, url, body)
}

// Post post请求
func (c *client) Post(ctx context.Context, url string, body any) (*httpclient.Response, error) {
	return c.Do(ctx, http.MethodPost, url, body)
}

// Put put请求
func (c *client) Put(ctx context.Context, url string, body any) (*httpclient.Response, error) {
	return c.Do(ctx, http.MethodPut, url, body)
}

// Delete delete请求
func (c *client) Delete(ctx context.Context, url string, body any) (*httpclient.Response, error) {
	return c.Do(ctx, http.MethodDelete, url, body)
}

// Do 做请求
func (c *client) Do(ctx context.Context, method string, url string, body any) (*httpclient.Response, error) {
	bodyBuffer, err := httpclient.HandleRequestBody(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyBuffer)
	if err != nil {
		return nil, err
	}

	interceptor := MetricsInterceptor(c.name, nil)
	req, handler := interceptor(req)

	resp, err := c.service.DoRequest(req)
	handler(resp, err)
	if err != nil {
		return nil, err
	}
	defer httpclient.Closeq(resp.Body)

	response := &httpclient.Response{
		RawResponse: resp,
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, fmt.Errorf("read response body error: %+v", err)
	}
	response.SetBody(respBody)
	response.SetSize(int64(len(response.Body())))

	return response, err
}
