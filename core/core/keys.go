package core

import (
	gnarkprecomputes "github.com/akakou/gnark-precomputes"
)

var UpdateRequestVerifyingKeys []gnarkprecomputes.PreparableCircuit

type KeyIndex struct {
	First, Second int64
}

type ProvingKey SnarkProver
type ProvingKeys map[RevocationListSize]SnarkProver

type VerifyingKey VerifyKey
type VerifyingKeys map[RevocationListSize]VerifyingKey
