package core

import (
	gnarkprecomputes "github.com/akakou/gnark-precomputes"
	gnarkserializable "github.com/akakou/gnark-serializable"
)

var UpdateRequestVerifyingKeys []gnarkprecomputes.PreparableCircuit

type KeyIndex struct {
	First, Second int64
}

type ProvingKeys []*SnarkProver

type VerifyingKeys []*gnarkserializable.VerifyingKey
type SizedVerifyingKeys []*SizedSnarkVerifier
