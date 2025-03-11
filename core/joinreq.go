package core

import (
	"encoding/json"

	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/highlevel"
)

type JoinRequest[T any] struct {
	Period        int64
	UserPublicKey []byte
	Proof         []byte
	Option        T
}

func RequestJoin[T any](prover []byte, option T) ([]byte, error) {
	period := utils.Today()

	proverObj := highlevel.HighLevelSnarkProver{}
	err := json.Unmarshal(prover, &proverObj)

	if err != nil {
		return nil, err
	}

	proof, reqObj, err := highlevel.JoinRequest(period, &proverObj)
	if err != nil {
		return nil, err
	}

	req := JoinRequest[T]{
		Period:        reqObj.Period,
		UserPublicKey: reqObj.UserPublicKey,
		Proof:         proof,
		Option:        option,
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return reqBytes, nil

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
