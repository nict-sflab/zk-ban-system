package gm

import (
	"errors"
	"fmt"
	"math"

	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban/precomputes"
	"github.com/akakou/zk-ban/snark"
	curve_bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	fr_bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	groth16_bls12381 "github.com/consensys/gnark/backend/groth16/bls12-381"
)

type PreparedUpdateSnarkVerifier struct {
	VerifierKey *precomputes.PreparedUpdateRequestVerifyingKey[fr_bls12381.Vector, *curve_bls12381.G1Jac, *groth16_bls12381.Proof]
	Prepared    *curve_bls12381.G1Jac
}

var errRLTooLarge = errors.New("RL is too large")

func SelectProperVerifier(verifierKeys core.SizedVerifyingKeys, rlSize *core.RevocationListSize) (*snark.SizedSnarkVerifier, error) {
	distance := math.MaxFloat64
	var verifyKey *snark.SizedSnarkVerifier = nil
	var err error = nil

	for _, vk := range verifierKeys {
		isFit := len(rlSize.NymsNumberPerPeriod) <= len(vk.RLSize)

		for i := range rlSize.NymsNumberPerPeriod {
			isFit = rlSize.NymsNumberPerPeriod[i] <= vk.RLSize[i] && isFit
		}

		if !isFit {
			continue
		}

		if verifierKeys == nil {
			verifyKey = vk
			continue
		}

		d := (&core.RevocationListSize{vk.RLSize}).Distance(core.RevocationListSizeWeightSetting)

		if distance > d {
			distance = d
			verifyKey = vk
		}
	}

	if verifyKey == nil {
		err = fmt.Errorf("%v: %v", errRLTooLarge, rlSize.NymsNumberPerPeriod)
	}

	return verifyKey, err
}
