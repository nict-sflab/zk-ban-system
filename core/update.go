package core

import (
	"encoding/json"

	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/highlevel"
	"github.com/akakou/zk-ban/witness"
)

type UpdateRequest struct {
	Before        int64
	After         int64
	UserPublicKey []byte
	Proof         []byte
}

func RequestUpdate(signer, rl, gpk, prover []byte) ([]byte, error) {
	today := utils.Today()

	proverObj := highlevel.HighLevelSnarkProver{}
	err := json.Unmarshal(prover, &proverObj)

	if err != nil {
		return nil, err
	}

	rlObj := witness.RevocationList{}
	err = json.Unmarshal(rl, &rlObj)

	if err != nil {
		return nil, err
	}

	signerObj := highlevel.HighLevelSigner{}
	err = json.Unmarshal(signer, &signerObj)

	if err != nil {
		return nil, err
	}

	updatedSigner, proof, err := highlevel.UpdateRequest(today, &signerObj, rlObj, gpk, &proverObj)
	if err != nil {
		return nil, err
	}

	req := UpdateRequest{
		After:         updatedSigner.Period,
		Before:        signerObj.Period,
		UserPublicKey: updatedSigner.UserPublicKey,
		Proof:         proof,
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return reqBytes, nil
}

func UpdateCredential(req, signer, gpk []byte, rl witness.RevocationList, verifyKey []byte) ([]byte, error) {
	period := utils.Today()

	var reqObj UpdateRequest
	err := json.Unmarshal(req, &reqObj)
	if err != nil {
		return nil, err
	}

	var signerObj highlevel.HighLevelSigner
	err = json.Unmarshal(signer, &signerObj)
	if err != nil {
		return nil, err
	}

	err = highlevel.VerifyUpdateRequest(reqObj.Proof, reqObj.UserPublicKey, period, &signerObj, rl, gpk, verifyKey)
	if err != nil {
		return nil, err
	}

	cred, err := json.Marshal(signer)
	if err != nil {
		return nil, err
	}

	return cred, nil

}
