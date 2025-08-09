package signer

import (
	"bytes"
	"fmt"
	"net/http"
	neturl "net/url"

	coresigner "github.com/akakou/zk-ban-system/core/signer"
)

func Sign(message []byte, count int64, signer, gpk []byte, url string) ([]byte, error) {
	signature, err := coresigner.Sign(message, count, signer, gpk)
	if err != nil {
		return nil, err
	}

	signatureStr := string(signature)

	parsed, err := neturl.Parse(url)
	if err != nil {
		return nil, err
	}

	query := parsed.Query()
	query.Add("signature", signatureStr)
	parsed.RawQuery = query.Encode()

	res, err := http.Get(parsed.String())
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code is %v, not 200\n%v", res.StatusCode, buf.String())
	}

	return buf.Bytes(), nil
}
