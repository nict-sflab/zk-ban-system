package core

import (
	gnarkprecomputes "github.com/akakou/gnark-precomputes"
)

var UpdateRequestVerifyingKeys []gnarkprecomputes.PreparableCircuit

type Witness = any
type Prepared = any
type Proof = any

type SnarkKey struct {
	VerifyingKey []byte
	Size         *RevocationListSize
}

var SnarkKeys map[KeyIndex]*SnarkKey = make(map[KeyIndex]*SnarkKey)

type KeyIndex struct {
	First, Second int64
}
