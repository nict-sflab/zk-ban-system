package core

import (
	gnarkprecomputes "github.com/akakou/gnark-precomputes"
)

var UpdateRequestVerifyingKeys []gnarkprecomputes.PreparableCircuit

type Witness = any
type Prepared = any
type Proof = any

type SnarkKey []byte

type KeyIndex struct {
	First, Second int64
}

type SnarkKeys map[RevocationListSize]*SnarkKey
