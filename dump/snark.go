package dump

import (
	"encoding/json"
	"math/big"

	gnarkserializable "github.com/akakou/gnark-serializable"
	"github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban/circuit"
	"github.com/akakou/zk-ban/precomputes"
	"github.com/akakou/zk-ban/snark"
	"github.com/akakou/zk-ban/witness"
	"github.com/consensys/gnark/frontend"
)

func Prepare[T frontend.Circuit](c T) ([]byte, []byte, error) {
	cc, err := snark.InitSNARK(c)
	if err != nil {
		return nil, nil, err
	}

	verifierBuf, err := json.Marshal(&gnarkserializable.VerifyingKey{cc.VerifyKey})
	if err != nil {
		return nil, nil, err
	}

	prover := core.SnarkProver{
		ConstraintSystem: &gnarkserializable.ConstraintSystem{cc.ConstraintSystem},
		ProveKey:         &gnarkserializable.ProvingKey{cc.ProveKey},
	}

	proverBuf, err := json.Marshal(&prover)
	if err != nil {
		return nil, nil, err
	}

	return proverBuf, verifierBuf, nil
}

func JoinRequestCircuit() ([]byte, []byte, error) {
	return Prepare(&circuit.JoinRequestCircuit{})
}

func SignCircuit() ([]byte, []byte, error) {
	return Prepare(&circuit.SignCircuit{})
}

func UpdateCircuit(nymsNumberPerSession, sessionNumber int) ([]byte, []byte, error) {
	rlWit := witness.RevocationList{}
	for range sessionNumber {
		rns := witness.RevokedNymsPerSession{
			SessionTag: big.NewInt(0),
			Nyms:       make([]*big.Int, nymsNumberPerSession),
		}

		rlWit = append(rlWit, rns)
	}

	return Prepare(&precomputes.UpdateCircuit{
		UpdateCircuit: &circuit.UpdateCircuit{
			RevocationList: circuit.NewRevocationListWitness(rlWit),
		},
	})
}
