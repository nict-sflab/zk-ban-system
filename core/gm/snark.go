package gm

import (
	"math"

	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban/precomputes"
	curve_bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	fr_bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	groth16_bls12381 "github.com/consensys/gnark/backend/groth16/bls12-381"
)

type PreparedSnarkVerifier struct {
	VerifierKey *precomputes.PreparedUpdateRequestVerifyingKey[fr_bls12381.Vector, *curve_bls12381.G1Jac, *groth16_bls12381.Proof]
	Prepared    *curve_bls12381.G1Jac
}

func SelectProperVerifier(verifierKeys core.SizedVerifyingKeys, rlSize *core.RevocationListSize) (*core.SizedSnarkVerifier, int) {
	distance := math.MaxFloat64
	var verifyKey *core.SizedSnarkVerifier = nil
	var result = -1

	for i, vk := range verifierKeys {
		isFit := len(rlSize.NymsNumberPerSession) <= len(vk.RLSize.NymsNumberPerSession)

		for i := range rlSize.NymsNumberPerSession {
			isFit = rlSize.NymsNumberPerSession[i] <= vk.RLSize.NymsNumberPerSession[i] && isFit
		}

		if !isFit {
			continue
		}

		if verifierKeys == nil {
			verifyKey = vk
			result = i
			continue
		}

		d := vk.RLSize.Distance(core.RevocationListSizeWeightSetting)

		if distance > d {
			distance = d
			verifyKey = vk
			result = i
		}
	}

	return verifyKey, result
}
