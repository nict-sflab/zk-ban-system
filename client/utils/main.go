package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
)

func FetchBinary(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		return nil, err
	}

	result := buf.Bytes()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code is %v, not 200\n%v\n%v", res.StatusCode, res, string(result))
	}

	return result, nil
}

func FetchBase64(url string) ([]byte, error) {
	buf, err := FetchBinary(url)
	if err != nil {
		return nil, err
	}

	return base64.URLEncoding.DecodeString(string(buf))
}
