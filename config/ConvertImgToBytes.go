package config

import (
	"bytes"
	"io"
	"mime/multipart"
)

func ConvertImgToBytes(image multipart.File) ([]byte, error) {
	buffer := &bytes.Buffer{}
	_, err := io.Copy(buffer, image)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
