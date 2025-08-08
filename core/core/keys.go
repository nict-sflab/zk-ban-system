package core

import (
	gnarkprecomputes "github.com/akakou/gnark-precomputes"
	"github.com/akakou/zk-ban/snark"
	"github.com/consensys/gnark/backend/groth16"
)

var UpdateRequestVerifyingKeys []gnarkprecomputes.PreparableCircuit

type KeyIndex struct {
	First, Second int64
}

type ProvingKey snark.SnarkProver
type ProvingKeys map[RevocationListSize]*ProvingKey

type VerifyingKey groth16.VerifyingKey
type VerifyingKeys map[RevocationListSize]groth16.VerifyingKey
