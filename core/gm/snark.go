package gm

import (
	gnarkprecomputes "github.com/akakou/gnark-precomputes"
	corecore "github.com/akakou/zk-ban-system/core/core"
)

var UpdateRequestVerifyingKeys []gnarkprecomputes.PreparableCircuit

type Witness = any
type Prepared = any
type Proof = any

type PreparableSnarkVerifierKey struct {
	VerifyingKey []byte
	Size         *corecore.RevocationListSize
}

var PreparableSnarkVerifierKeys = []*PreparableSnarkVerifierKey{}

type PreparedSnarkVerifier struct {
	VerifierKey *PreparableSnarkVerifierKey
	Prepared    []byte
}

type DoubleMapKey struct {
	First, Second int
}

var PreparedSnarkVerifiers map[DoubleMapKey]*PreparedSnarkVerifier = make(map[DoubleMapKey]*PreparedSnarkVerifier)
