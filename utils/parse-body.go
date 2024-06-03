package utils

import (
	"encoding/json"
	"io"
)

func ParseBody(body io.ReadCloser, result interface{}) error {
	defer body.Close()
	decoder := json.NewDecoder(body)
	return decoder.Decode(result)
}