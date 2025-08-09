package gm

import (
	"math"

	corecore "github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban/precomputes"
	curve_bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	fr_bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	groth16_bls12381 "github.com/consensys/gnark/backend/groth16/bls12-381"
)

type PreparedSnarkVerifier struct {
	VerifierKey *precomputes.PreparedUpdateRequestVerifyingKey[fr_bls12381.Vector, *curve_bls12381.G1Jac, *groth16_bls12381.Proof]
	Prepared    *curve_bls12381.G1Jac
}

func QueryProperRLWitSize(verifierKeys corecore.VerifyingKeys, size *corecore.RevocationListSize) *corecore.RevocationListSize {
	distance := math.MaxFloat64
	var result *corecore.RevocationListSize = nil

	for i := range verifierKeys {
		isNotFitSize1 := i.NymsNumberPerSession < size.NymsNumberPerSession
		isNotFitSize2 := i.SessionNumber < size.SessionNumber

		if isNotFitSize1 || isNotFitSize2 {
			continue
		}

		if verifierKeys == nil {
			result = &i
			continue
		}

		d := i.Distance(corecore.RevocationListSizeWeightSetting)

		if distance > d {
			distance = d
			result = &i
		}
	}

	return result
}
