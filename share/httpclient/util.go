package httpclient

import "io"

func Closeq(v interface{}) {
	if c, ok := v.(io.Closer); ok {
		silently(c.Close())
	}
}

func silently(_ ...interface{}) {}
