//go:build !js

package moreio

import (
	"context"
	"os"
	"path/filepath"
)

func Open(ctx context.Context, path string) (File, error) {
	return os.Open(filepath.FromSlash(path))
}
