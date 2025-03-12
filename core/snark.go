package core

import (
	"encoding/json"
	"math/big"

	snark "github.com/akakou/snark-utils"
	snarkencode "github.com/akakou/snark-utils/encode"
	"github.com/akakou/zk-ban/circuit"
	"github.com/akakou/zk-ban/witness"
	"github.com/consensys/gnark/frontend"
)

func Prepare[T frontend.Circuit](c T) ([]byte, []byte, error) {
	cc, err := snark.InitSNARK(c)
	if err != nil {
		return nil, nil, err
	}

	encodedProveKey, err := snarkencode.EncodeProverKey(cc.ProveKey)
	if err != nil {
		return nil, nil, err
	}

	encodedVerifyKey, err := snarkencode.EncodeVerifierKey(cc.VerifyKey)
	if err != nil {
		return nil, nil, err
	}

	encodedCircuit, err := snarkencode.EncodeCircuit(cc.ConstraintSystem)
	if err != nil {
		return nil, nil, err
	}

	prover := snarkencode.HighLevelSnarkProver{
		ConstraintSystem: encodedCircuit,
		ProveKey:         encodedProveKey,
	}

	proverBuf, err := json.Marshal(&prover)
	if err != nil {
		return nil, nil, err
	}

	return proverBuf, encodedVerifyKey, nil
}

func JoinRequestCircuit() ([]byte, []byte, error) {
	return Prepare(&circuit.JoinRequestCircuit{})
}

func SignCircuit() ([]byte, []byte, error) {
	return Prepare(&circuit.SignCircuit{})
}

func UpdateCircuit(rl []int32) ([]byte, []byte, error) {
	rlWit := witness.RevocationList{}
	for _, r := range rl {
		rns := witness.RevokedNymsPerSession{
			SessionTag: big.NewInt(0),
			Nyms:       make([]*big.Int, r),
		}

		rlWit = append(rlWit, rns)
	}

	return Prepare(&circuit.UpdateCircuit{
		RevocationList: circuit.NewRevocationListWitness(rlWit),
	})
}
