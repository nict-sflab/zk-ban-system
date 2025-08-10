package signer

import (
	"bytes"
	"encoding/base64"
	"net/http"

	coresigner "github.com/akakou/zk-ban-system/core/signer"
	"github.com/akakou/zk-ban-system/utils"
)

func RequestUpdate(signer, rl, gpk []byte, url string) ([]byte, error) {
	now := utils.Period()

	req, err := coresigner.RequestUpdate(rl, signer, now, gpk)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(req))
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

	newSignerFinalized, err := coresigner.SetCredential(cred, now, signer)
	if err != nil {
		return nil, err
	}

	return newSignerFinalized, nil
}
