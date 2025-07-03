package signer

import (
	"encoding/json"
	"math/big"

	"github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/dump"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/highlevel"
)

func RequestJoin[T any](option T) ([]byte, []byte, error) {
	period := utils.Today()

	proverObj := highlevel.HighLevelSnarkProver{}
	err := json.Unmarshal(dump.JoinProverKey, &proverObj)

	if err != nil {
		return nil, nil, err
	}

	proof, reqObj, assign, err := highlevel.JoinRequest(period, &proverObj)
	if err != nil {
		return nil, nil, err
	}

	signer := highlevel.HighLevelSigner{
		Secret:        assign.UserSecretKey.(*big.Int).Bytes(),
		UserPublicKey: reqObj.UserPublicKey,
		Period:        period,
	}

	signerBytes, err := json.Marshal(signer)
	if err != nil {
		return nil, nil, err
	}

	req := core.JoinRequest[T]{
		Period:        reqObj.Period,
		UserPublicKey: reqObj.UserPublicKey,
		Proof:         proof,
		Option:        option,
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, nil, err
	}

	return reqBytes, signerBytes, nil

}

func SetCredential(cred []byte, signer []byte) ([]byte, error) {
	signerObj := highlevel.HighLevelSigner{}
	err := json.Unmarshal(signer, &signerObj)

	if err != nil {
		return nil, err
	}

	signerObj.Credential = cred

	signerBytes, err := json.Marshal(signerObj)
	if err != nil {
		return nil, err
	}

	return signerBytes, nil
}
