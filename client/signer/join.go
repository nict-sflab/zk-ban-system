package signer

import (
	"bytes"
	"encoding/base64"
	"net/http"

	coresigner "github.com/akakou/zk-ban-system/core/signer"
)

func RequestJoin(idToken string, prover []byte, url string) ([]byte, error) {
	requestBytes, signer, err := coresigner.RequestJoin(prover, idToken)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(requestBytes))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		return nil, err
	}

	cred, err := base64.URLEncoding.DecodeString(buf.String())
	if err != nil {
		return nil, err
	}

	signer, err = coresigner.SetCredential(cred, signer)
	if err != nil {
		return nil, err
	}

	return signer, nil
}
