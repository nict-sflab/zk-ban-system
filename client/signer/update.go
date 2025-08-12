package signer

import (
	signerutils "github.com/akakou/zk-ban-system/client/utils"
	coresigner "github.com/akakou/zk-ban-system/core/signer"
	"github.com/akakou/zk-ban-system/utils"
)

func RequestUpdate(signer, rl, gpk []byte, url string) ([]byte, error) {
	now := utils.Period()

	req, err := coresigner.RequestUpdate(rl, signer, now, gpk)
	if err != nil {
		return nil, err
	}

	cred, err := signerutils.FetchBase64WithPOST(req, url)
	if err != nil {
		return nil, err
	}

	newSignerFinalized, err := coresigner.SetCredential(cred, now, signer)
	if err != nil {
		return nil, err
	}

	return newSignerFinalized, nil
}
