package signer

import (
	zkban "github.com/akakou/zk-ban"
	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/utils/codec"
	"github.com/akakou/zk-ban/load"
	zkbanw "github.com/akakou/zk-ban/witness"
)

func RequestJoin[T any](now int64, option T) ([]byte, []byte, error) {
	proverObj, err := load.LoadUserBasicKey("join")
	if err != nil {
		return nil, nil, err
	}

	proof, usk, err := zkban.RequestJoin(now, proverObj)
	if err != nil {
		return nil, nil, err
	}

	signer := zkbanw.Signer{
		UserSecretKey: usk,
		Credential: &zkbanw.Credential{
			Signature: []byte{0},
		},
		Period: now,
	}

	signerBytes, err := codec.Marshal(signer)
	if err != nil {
		return nil, nil, err
	}

	req := core.JoinRequest[T]{
		Period:      now,
		JoinRequest: *proof,
		Option:      option,
	}

	reqBytes, err := codec.Marshal(req)
	if err != nil {
		return nil, nil, err
	}

	return reqBytes, signerBytes, nil

}

func SetCredential(cred []byte, period int64, signer []byte) ([]byte, error) {
	signerObj := zkbanw.Signer{}
	err := codec.Unmarshal(signer, &signerObj)

	if err != nil {
		return nil, err
	}

	signerObj.Credential.Signature = cred
	signerObj.Period = period

	signerBytes, err := codec.Marshal(signerObj)
	if err != nil {
		return nil, err
	}

	return signerBytes, nil
}
