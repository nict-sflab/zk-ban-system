package core

import (
	gnarkserializable "github.com/akakou/gnark-serializable"
	"github.com/akakou/zk-ban/snark"
)

type SnarkProver struct {
	ConstraintSystem *gnarkserializable.ConstraintSystem
	ProveKey         *gnarkserializable.ProvingKey
}

func (prover *SnarkProver) CoreKey() *snark.SnarkProver {
	return &snark.SnarkProver{
		ConstraintSystem: prover.ConstraintSystem.ConstraintSystem,
		ProveKey:         prover.ProveKey.ProvingKey,
	}
}

type SizedSnarkVerifier struct {
	VerifyKey *gnarkserializable.VerifyingKey
	RLSize    *RevocationListSize
}
