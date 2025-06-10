package signer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	coresigner "github.com/akakou/zk-ban-system/core/signer"
)

func Sign(message []byte, count int64, signer, gpk, prover []byte, url string) ([]byte, error) {
	signature, err := coresigner.Sign(message, count, signer, gpk, prover)
	if err != nil {
		return nil, err
	}

	signatureBuf, err := json.Marshal(signature)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(signatureBuf))
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
