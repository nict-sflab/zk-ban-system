package core

import (
	gnarkprecomputes "github.com/akakou/gnark-precomputes"
	gnarkserializable "github.com/akakou/gnark-serializable"
	"github.com/akakou/zk-ban/snark"
)

var UpdateRequestVerifyingKeys []gnarkprecomputes.PreparableCircuit

type KeyIndex struct {
	First, Second int64
}

type ProvingKeys []*snark.SnarkProver

type VerifyingKeys []*gnarkserializable.VerifyingKey
type SizedVerifyingKeys []*snark.SizedSnarkVerifier
