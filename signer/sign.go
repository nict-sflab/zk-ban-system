package signer

import (
	"encoding/json"

	"github.com/akakou/zk-ban-system/core"
)

func Sign(message []byte, count int64, signer []byte, gpk []byte, prover []byte) ([]byte, error) {
	signature, err := core.Sign(message, count, signer, gpk, prover)
	if err != nil {
		return nil, err
	}

	sigBytes, err := json.Marshal(signature)
	if err != nil {
		return nil, err
	}

	return sigBytes, err
}
