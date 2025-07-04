package signer

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/akakou/zk-ban/highlevel"
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

func FetchRevocationList(signer []byte, u string) ([]byte, error) {
	signerObj := highlevel.HighLevelSigner{}
	err := json.Unmarshal(signer, &signerObj)
	if err != nil {
		return nil, err
	}

	uu, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	query := uu.Query()
	query.Add("before", fmt.Sprintf("%d", signerObj.Period))
	uu.RawQuery = query.Encode()

	res, err := http.Get(uu.String())
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

	return buf.Bytes(), nil
}
