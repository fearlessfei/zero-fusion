package httpc

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpc"
)

type RoundTripper struct {
}

func (r *RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := httpc.DoRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
