package signer

import (
	zkban "github.com/akakou/zk-ban"
	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/utils/codec"
	"github.com/akakou/zk-ban/load"
	"github.com/akakou/zk-ban/primitives"
	zkbanw "github.com/akakou/zk-ban/witness"
)

func Sign(message []byte, count int64, signer []byte, gpk []byte) ([]byte, error) {
	prover, err := load.LoadUserBasicKey("sign")
	if err != nil {
		return nil, err
	}

	signerObj := zkbanw.Signer{}
	err = codec.Unmarshal(signer, &signerObj)
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
		prover,
	)
	if err != nil {
		return nil, err
	}

	signatureObj := &core.Signature{
		Signature: signature,
		Message:   message,
		Count:     count,
	}

	signatureBuf, err := codec.Marshal(signatureObj)
	if err != nil {
		return nil, err
	}

	return signatureBuf, nil
}
