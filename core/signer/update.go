package signer

import (
	gnarkserializable "github.com/akakou/gnark-serializable"
	zkban "github.com/akakou/zk-ban"
	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban-system/utils/codec"
	"github.com/akakou/zk-ban/load"
	"github.com/akakou/zk-ban/primitives"
	"github.com/akakou/zk-ban/snark"
	zkbanw "github.com/akakou/zk-ban/witness"
	"github.com/consensys/gnark/backend/groth16"
)

func RequestUpdate(rl, signer []byte, now int64, gpk []byte) ([]byte, error) {
	var rlObj core.RevocationList
	err := codec.Unmarshal(rl, &rlObj)
	if err != nil {
		return nil, err
	}

	prover, err := load.LoadUserKey(rlObj.KeyName, "update")
	if err != nil {
		return nil, err
	}

	var signerObj zkbanw.Signer
	err = codec.Unmarshal(signer, &signerObj)
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

	reqBytes, err := codec.Marshal(req)
	if err != nil {
		return nil, err
	}

	return reqBytes, nil
}

func RequestCheckUpdateIsExist(signer []byte, now int64) ([]byte, error) {
	var signerObj zkbanw.Signer
	err := codec.Unmarshal(signer, &signerObj)
	if err != nil {
		return nil, err
	}

	nextPeriod := utils.Period()
	nextPublicKey, err := signerObj.UserSecretKey.PublicKey(nextPeriod)
	if err != nil {
		return nil, err
	}

	req := core.UpdateRequest{
		Before: signerObj.Period,
		UpdateRequest: &zkban.UpdateRequest{
			PublicKey: nextPublicKey,
			UpdateTicket: &zkbanw.OneTimeTicket{
				Number: primitives.NewBigInt(0),
			},
			Proof: gnarkserializable.Proof{
				Proof: groth16.NewProof(snark.EcCurve),
			},
		},
	}

	reqBytes, err := codec.Marshal(req)
	if err != nil {
		return nil, err
	}

	return reqBytes, nil
}
