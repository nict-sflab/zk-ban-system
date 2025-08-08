package signer

import (
	"encoding/json"

	zkban "github.com/akakou/zk-ban"
	"github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/load"
	"github.com/akakou/zk-ban-system/utils"
	zkbanw "github.com/akakou/zk-ban/witness"
)

func RequestJoin[T any](option T) ([]byte, []byte, error) {
	period := utils.Today()

	proverObj := core.SnarkProver{}
	err := json.Unmarshal(load.JoinProverKey, &proverObj)

	if err != nil {
		return nil, nil, err
	}

	proof, usk, err := zkban.RequestJoin(period, proverObj.CoreKey())
	if err != nil {
		return nil, nil, err
	}

	signer := zkbanw.Signer{
		UserSecretKey: usk,
		Credential:    nil,
		Period:        period,
	}

	signerBytes, err := json.Marshal(signer)
	if err != nil {
		return nil, nil, err
	}

	req := core.JoinRequest[T]{
		Period:      period,
		JoinRequest: proof,
		Option:      option,
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, nil, err
	}

	return reqBytes, signerBytes, nil

}

func SetCredential(cred []byte, signer []byte) ([]byte, error) {
	signerObj := zkbanw.Signer{}
	err := json.Unmarshal(signer, &signerObj)

	if err != nil {
		return nil, err
	}

	signerObj.Credential.Signature = cred

	signerBytes, err := json.Marshal(signerObj)
	if err != nil {
		return nil, err
	}

	return signerBytes, nil
}
