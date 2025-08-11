package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
)

func FetchBinaryWithPOST(request []byte, url string) ([]byte, error) {
	res, err := http.Post(url, "application/json", bytes.NewBuffer(request))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func FetchBase64WithPOST(request []byte, url string) ([]byte, error) {
	buf, err := FetchBinaryWithPOST(request, url)
	if err != nil {
		return nil, err
	}

	result, err := base64.URLEncoding.DecodeString(string(buf))
	if err != nil {
		return nil, fmt.Errorf("%e binary is %s", err, buf)
	}

	return result, nil
}
