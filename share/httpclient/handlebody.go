package httpclient

import (
	"bytes"
	"errors"
	"io"
	"reflect"

	jsoniter "github.com/json-iterator/go"

	"zero-fusion/share/reflectx"
)

func HandleRequestBody(body any) (*bytes.Buffer, error) {
	var bodyBytes []byte
	kind := reflectx.KindOf(body)

	var (
		bodyBuf *bytes.Buffer
		err     error
	)

	if reader, ok := body.(io.Reader); ok {
		bodyBuf = &bytes.Buffer{}
		_, err = bodyBuf.ReadFrom(reader)
		if err != nil {
			return nil, err
		}
	} else if b, ok := body.([]byte); ok {
		bodyBytes = b
	} else if s, ok := body.(string); ok {
		bodyBytes = []byte(s)
	} else if kind == reflect.Struct || kind == reflect.Map || kind == reflect.Slice {
		bodyBytes, err = jsoniter.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	if bodyBytes == nil && bodyBuf == nil {
		err = errors.New("unsupported 'Body' type/value")
		if err != nil {
			return nil, err
		}
	}

	// []byte into Buffer
	if bodyBytes != nil {
		bodyBuf = &bytes.Buffer{}
		_, _ = bodyBuf.Write(bodyBytes)
	}

	return bodyBuf, nil
}
