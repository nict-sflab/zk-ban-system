package signer

import (
	"encoding/json"

	zkban "github.com/akakou/zk-ban"
	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/keys"
	zkbanw "github.com/akakou/zk-ban/witness"
)

func RequestUpdate(rl, signer []byte, now int64, gpk []byte) ([]byte, error) {
	var rlObj core.RevocationList
	err := json.Unmarshal(rl, &rlObj)
	if err != nil {
		return nil, err
	}

	provers, err := keys.LoadUserUpdateKey(keys.UpdateProverKey)
	if err != nil {
		return nil, err
	}

	prover := provers[rlObj.Index]

	var signerObj zkbanw.Signer
	err = json.Unmarshal(signer, &signerObj)
	if err != nil {
		return nil, err
	}

	gpkObj, err := zkbanw.GroupPublicKeyFromBytes(gpk)
	if err != nil {
		return nil, err
	}

	coreReq, err := zkban.RequestUpdate(now, &signerObj, *rlObj.List, gpkObj, prover)
	if err != nil {
		return nil, err
	}

	req := core.UpdateRequest{
		Before:        signerObj.Period,
		UpdateRequest: coreReq,
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return reqBytes, nil
}
