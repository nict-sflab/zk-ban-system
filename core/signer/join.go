package signer

import (
	"encoding/json"
	"fmt"

	"github.com/akakou/snark-utils/encode"
	zkban "github.com/akakou/zk-ban"
	"github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/load"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/snark"
	zkbanw "github.com/akakou/zk-ban/witness"
)

func RequestJoin[T any](option T) ([]byte, []byte, error) {
	period := utils.Today()

	proverObj := encode.HighLevelSnarkProver{}
	err := json.Unmarshal(load.JoinProverKey, &proverObj)

	if err != nil {
		return nil, nil, err
	}

	// todo:
	tmp, err := proverObj.ToSnarkProver()
	if err != nil {
		return nil, nil, err
	}

	prover := snark.SnarkProver{
		ConstraintSystem: tmp.ConstraintSystem,
		ProveKey:         tmp.ProveKey,
	}

	proof, usk, err := zkban.RequestJoin(period, &prover)
	if err != nil {
		return nil, nil, err
	}

	signer := zkbanw.Signer{
		UserSecretKey: usk,
		Credential: &zkbanw.Credential{
			Signature: []byte{0},
		},
		Period: period,
	}

	signerBytes, err := json.Marshal(signer)
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("aaa %v\n", string(signerBytes))

	req := core.JoinRequest[T]{
		Period:      period,
		JoinRequest: *proof,
		Option:      option,
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, nil, err
	}
	fmt.Printf("bbb %v\n", string(reqBytes))
	fmt.Printf("bbb %x\n", req.JoinRequest.UserPublicKey)

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
