package signer

import (
	"encoding/json"

	"github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban/highlevel"
)

func Sign(message []byte, count int64, signer []byte, gpk []byte, prover []byte) (*core.Signature, error) {
	proverObj := highlevel.HighLevelSnarkProver{}
	err := json.Unmarshal(prover, &proverObj)

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

	// sigBytes, err := json.Marshal(signature)
	// if err != nil {
	// 	return nil, err
	// }

	return &core.Signature{
		Signature: signature,
		Message:   message,
	}, err
}
