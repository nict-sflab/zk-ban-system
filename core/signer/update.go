package signer

import (
	"encoding/json"

	"github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/load"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/highlevel"
)

func RequestUpdate(rl, signer, gpk []byte) ([]byte, []byte, error) {
	today := utils.Today()

	var rlObj core.RevocationList
	err := json.Unmarshal(rl, &rlObj)
	if err != nil {
		return nil, nil, err
	}

	provers, err := load.LoadUserUpdateKey()
	if err != nil {
		return nil, nil, err
	}

	prover := provers[*rlObj.Size]

	var proverObj highlevel.HighLevelSnarkProver
	err = json.Unmarshal(*prover, &proverObj)
	if err != nil {
		return nil, nil, err
	}

	var signerObj highlevel.HighLevelSigner
	err = json.Unmarshal(signer, &signerObj)

	if err != nil {
		return nil, nil, err
	}

	updatedSigner, proof, err := highlevel.UpdateRequest(today, &signerObj, *rlObj.List, gpk, &proverObj)
	if err != nil {
		return nil, nil, err
	}

	req := core.UpdateRequest{
		After:         updatedSigner.Period,
		Before:        signerObj.Period,
		UserPublicKey: updatedSigner.UserPublicKey,
		Proof:         proof,
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, nil, err
	}

	updatedSignerBytes, err := json.Marshal(updatedSigner)
	if err != nil {
		return nil, nil, err
	}

	return reqBytes, updatedSignerBytes, nil
}

// func UpdateCredential(upk, cred []byte, period int64, signer []byte) ([]byte, []byte, error) {
// 	signerObj := highlevel.HighLevelSigner{}
// 	err := json.Unmarshal(signer, &signerObj)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	signerObj.UserPublicKey = upk
// 	signerObj.Credential = cred
// 	signerObj.Period = period

// 	signerBytes, err := json.Marshal(signerObj)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	return signerBytes, nil
// }
