//go:build js

package moreio

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/context/ctxhttp"
)

type file struct {
	*bytes.Reader
}

func (f *file) Close() error {
	return nil
}

func Open(ctx context.Context, path string) (io.ReadSeekCloser, error) {
	res, err := ctxhttp.Get(ctx, http.DefaultClient, path)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return &file{bytes.NewReader(body)}, nil
}
