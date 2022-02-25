//go:build !js

package moreio

import (
	"context"
	"io"
	"os"
	"path/filepath"
)

func Open(ctx context.Context, path string) (io.ReadSeekCloser, error) {
	return os.Open(filepath.FromSlash(path))
}
