package routes

import (
	"io"
	"bytes"
)

func GetPayloadBytes(closer io.ReadCloser)([]byte ){
	buf := new(bytes.Buffer)
	buf.ReadFrom(closer)
	return buf.Bytes()
}
