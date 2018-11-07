package tools

import (
	"bytes"
	"io"
	"io/ioutil"
)

type ClosingBuffer struct {
	*bytes.Buffer
}

func (cb *ClosingBuffer) Close() error {
	return nil
}

func ReadAll(rc io.ReadCloser) ([]byte, error) {
	defer rc.Close()

	if cb, ok := rc.(*ClosingBuffer); ok {
		return cb.Bytes(), nil
	}

	return ioutil.ReadAll(rc)
}
