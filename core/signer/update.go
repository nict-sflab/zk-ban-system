package signer

import (
	"encoding/json"

	zkban "github.com/akakou/zk-ban"
	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/load"
	zkbanw "github.com/akakou/zk-ban/witness"
)

func RequestUpdate(rl, signer []byte, now int64, gpk []byte) ([]byte, error) {
	var rlObj core.RevocationList
	err := json.Unmarshal(rl, &rlObj)
	if err != nil {
		return nil, err
	}

	provers, err := load.LoadUserUpdateKey()
	if err != nil {
		return nil, err
	}

	prover := provers[*rlObj.Size]

	var signerObj zkbanw.Signer
	err = json.Unmarshal(signer, &signerObj)
	if err != nil {
		return nil, err
	}

	gpkObj, err := zkbanw.GroupPublicKeyFromBytes(gpk)
	if err != nil {
		return nil, err
	}

	coreReq, err := zkban.RequestUpdate(now, &signerObj, *rlObj.List, gpkObj, prover.CoreKey())
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
