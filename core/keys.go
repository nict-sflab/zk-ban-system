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

type ProvingKeys []*snark.SnarkProver

type VerifyingKeys []*groth16.VerifyingKey
type SizedVerifyingKeys []*snark.SizedSnarkVerifier
