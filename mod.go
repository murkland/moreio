package moreio

import "io"

type File interface {
	io.ReaderAt
	io.ReadSeekCloser
}
