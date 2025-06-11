package signer

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
)

func FetchGroupPublicKey(url string) ([]byte, error) {
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

	text := buf.String()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code is %v, not 200\n%v", res.StatusCode, text)
	}

	gpk, err := base64.URLEncoding.DecodeString(text)
	if err != nil {
		return nil, err
	}

	return gpk, nil
}
