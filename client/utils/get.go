package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
)

func FetchBinaryWithGET(url string) ([]byte, error) {
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

func FetchBase64WithGET(url string) ([]byte, error) {
	buf, err := FetchBinaryWithGET(url)
	if err != nil {
		return nil, err
	}

	result, err := base64.URLEncoding.DecodeString(string(buf))
	if err != nil {
		return nil, fmt.Errorf("%s\nmessage is `%s`", err.Error(), buf)
	}

	return result, nil
}
