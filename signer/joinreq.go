package signer

import (
	"bytes"
	"encoding/base64"
	"net/http"

	"github.com/akakou/zk-ban-system/core"
)

func RequestJoin(idToken []byte, prover []byte, url string) ([]byte, error) {
	requestBytes, signer, err := core.RequestJoin(prover, idToken)
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

	signer, err = core.SetCredential(cred, signer)
	if err != nil {
		return nil, err
	}

	return signer, nil
}
