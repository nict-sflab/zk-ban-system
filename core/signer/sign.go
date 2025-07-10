package signer

import (
	"encoding/json"

	"github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/load"
	"github.com/akakou/zk-ban/highlevel"
)

func Sign(message []byte, count int64, signer []byte, gpk []byte) ([]byte, error) {
	proverObj := highlevel.HighLevelSnarkProver{}
	err := json.Unmarshal(load.SignProverKey, &proverObj)

	if err != nil {
		return nil, err
	}

	signerObj := highlevel.HighLevelSigner{}
	err = json.Unmarshal(signer, &signerObj)

	if err != nil {
		return nil, err
	}

	signature, err := highlevel.Sign(
		message,
		count,
		signerObj,
		gpk,
		&proverObj,
	)
	if err != nil {
		return nil, err
	}

	signatureObj := &core.Signature{
		Signature: signature,
		Message:   message,
	}

	signatureBuf, err := json.Marshal(signatureObj)
	if err != nil {
		return nil, err
	}

	return signatureBuf, nil
}
