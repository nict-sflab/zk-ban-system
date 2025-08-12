package signer

import (
	signerutils "github.com/akakou/zk-ban-system/client/utils"
	coresigner "github.com/akakou/zk-ban-system/core/signer"
	"github.com/akakou/zk-ban-system/utils"
)

func RequestJoin(idToken string, url string) ([]byte, error) {
	now := utils.Period()
	requestBytes, signer, err := coresigner.RequestJoin(now, idToken)
	if err != nil {
		return nil, err
	}

	cred, err := signerutils.FetchBase64WithPOST(requestBytes, url)
	if err != nil {
		return nil, err
	}

	signer, err = coresigner.SetCredential(cred, now, signer)
	if err != nil {
		return nil, err
	}

	return signer, nil
}
