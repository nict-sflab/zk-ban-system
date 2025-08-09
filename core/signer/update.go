package signer

import (
	"encoding/json"

	zkban "github.com/akakou/zk-ban"
	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/load"
	"github.com/akakou/zk-ban-system/utils"
	zkbanw "github.com/akakou/zk-ban/witness"
)

func RequestUpdate(rl, signer, gpk []byte) ([]byte, []byte, error) {
	today := utils.Period()

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

	var signerObj zkbanw.Signer
	err = json.Unmarshal(signer, &signerObj)
	if err != nil {
		return nil, nil, err
	}

	gpkObj, err := zkbanw.GroupPublicKeyFromBytes(gpk)
	if err != nil {
		return nil, nil, err
	}

	coreReq, err := zkban.RequestUpdate(today, &signerObj, *rlObj.List, gpkObj, prover.CoreKey())
	if err != nil {
		return nil, nil, err
	}

	req := core.UpdateRequest{
		Before:        signerObj.Period,
		UpdateRequest: coreReq,
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, nil, err
	}

	signerObj.Period = today

	updatedSignerBytes, err := json.Marshal(signerObj)
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
