package core

import (
	"encoding/json"

	"github.com/akakou/zk-ban/highlevel"
)

type Signature struct {
	Signature *highlevel.Signature
	Message   []byte
}

func Sign(message []byte, count int64, signer []byte, gpk []byte, prover []byte) (*Signature, error) {
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

	return &Signature{
		Signature: signature,
		Message:   message,
	}, err
}

func Verify(signature *Signature, gpk, verifyKey []byte) error {
	err := highlevel.Verify(signature.Signature, signature.Message, gpk, verifyKey)
	if err != nil {
		return err
	}

	return nil
}
