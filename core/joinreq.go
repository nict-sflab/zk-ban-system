package core

import (
	"encoding/json"
	"math/big"

	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/highlevel"
)

type JoinRequest[T any] struct {
	Period        int64
	UserPublicKey []byte
	Proof         []byte
	Option        T
}

func RequestJoin[T any](prover []byte, option T) ([]byte, []byte, error) {
	period := utils.Today()

	proverObj := highlevel.HighLevelSnarkProver{}
	err := json.Unmarshal(prover, &proverObj)

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

	req := JoinRequest[T]{
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

func IssueCredential[T any](req, gsk, verifyKey []byte) ([]byte, error) {
	period := utils.Today()

	var reqObj JoinRequest[T]

	err := json.Unmarshal(req, &reqObj)
	if err != nil {
		return nil, err
	}

	err = highlevel.VerifyJoinReq(reqObj.Proof, reqObj.UserPublicKey, period, verifyKey)
	if err != nil {
		return nil, err
	}

	cred, err := highlevel.IssueCredential(period, reqObj.UserPublicKey, gsk)
	if err != nil {
		return nil, err
	}

	return cred, nil
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
