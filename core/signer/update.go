package signer

import (
	"encoding/json"

	"github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/highlevel"
)

func RequestUpdate(rl, signer, gpk, prover []byte) (*core.UpdateRequest, error) {
	today := utils.Today()

	proverObj := highlevel.HighLevelSnarkProver{}
	err := json.Unmarshal(prover, &proverObj)

	if err != nil {
		return nil, err
	}

	var rlObj core.RevocationList
	err = json.Unmarshal(rl, &rlObj)
	if err != nil {
		return nil, err
	}

	signerObj := highlevel.HighLevelSigner{}
	err = json.Unmarshal(signer, &signerObj)

	if err != nil {
		return nil, err
	}

	updatedSigner, proof, err := highlevel.UpdateRequest(today, &signerObj, *rlObj.List, gpk, &proverObj)
	if err != nil {
		return nil, err
	}

	req := core.UpdateRequest{
		After:         updatedSigner.Period,
		Before:        signerObj.Period,
		UserPublicKey: updatedSigner.UserPublicKey,
		Proof:         proof,
	}

	return &req, nil
}
