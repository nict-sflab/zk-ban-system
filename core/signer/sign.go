package signer

import (
	"encoding/json"

	zkban "github.com/akakou/zk-ban"
	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/keys"
	"github.com/akakou/zk-ban/primitives"
	"github.com/akakou/zk-ban/snark"
	zkbanw "github.com/akakou/zk-ban/witness"
)

func Sign(message []byte, count int64, signer []byte, gpk []byte) ([]byte, error) {
	proverObj := snark.SnarkProver{}
	err := json.Unmarshal(keys.SignProverKey, &proverObj)
	if err != nil {
		return nil, err
	}

	signerObj := zkbanw.Signer{}
	err = json.Unmarshal(signer, &signerObj)
	if err != nil {
		return nil, err
	}

	_, gpkObj, _ := zkbanw.RandomGroupKeyPair()
	_, err = gpkObj.SetBytes(gpk)
	if err != nil {
		return nil, err
	}

	signature, err := zkban.Sign(
		primitives.BigIntFromBytes(message),
		count,
		&signerObj,
		&zkbanw.GroupPublicKey{PublicKey: gpkObj},
		&proverObj,
	)
	if err != nil {
		return nil, err
	}

	signatureObj := &core.Signature{
		Signature: signature,
		Message:   message,
		Count:     count,
	}

	signatureBuf, err := json.Marshal(signatureObj)
	if err != nil {
		return nil, err
	}

	return signatureBuf, nil
}
