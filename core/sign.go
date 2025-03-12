package core

import (
	"encoding/json"

	snarkencode "github.com/akakou/snark-utils/encode"
	"github.com/akakou/zk-ban/highlevel"
)

func Sign(message []byte, count int64, signer []byte, gpk []byte, prover []byte) ([]byte, error) {
	proverObj := snarkencode.HighLevelSnarkProver{}
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

	sigBytes, err := json.Marshal(signature)
	if err != nil {
		return nil, err
	}

	return sigBytes, err
}

func Verify(signature, m, gpk, verifyKey []byte) error {
	sigObj := highlevel.Signature{}

	err := json.Unmarshal(signature, &sigObj)
	if err != nil {
		return err
	}

	err = highlevel.Verify(&sigObj, m, gpk, verifyKey)
	if err != nil {
		return err
	}

	return nil
}
