package gm

import (
	gnarkprecomputes "github.com/akakou/gnark-precomputes"
)

var UpdateRequestVerifyingKeys []gnarkprecomputes.PreparableCircuit

type Witness = any
type Prepared = any
type Proof = any

type PreparableSnarkVerifierKey struct {
	VerifyingKey []byte
	Size         *RevocationListSize
}

type PreparedSnarkVerifier struct {
	VerifierKey PreparableSnarkVerifierKey
	Prepared    []byte
}

type DoubleMapKey struct {
	First, Second int
}

var PreparedSnarkVerifiers map[DoubleMapKey]*PreparedSnarkVerifier = make(map[DoubleMapKey]*PreparedSnarkVerifier)
