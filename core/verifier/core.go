package verifier

import (
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/precomputes"
	zkbanw "github.com/akakou/zk-ban/witness"
	curve_bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"
	fr_bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381/fr"
	"github.com/consensys/gnark/backend/groth16"
	groth16_bls12381 "github.com/consensys/gnark/backend/groth16/bls12-381"
)

type Verifier struct {
	GroupPublicKey       *zkbanw.GroupPublicKey
	VerifyingKey         *groth16.VerifyingKey
	PreparedVerifyingKey *PreparedAuthSnarkVerifier
	Period               int64
}

type PreparedAuthSnarkVerifier struct {
	VerifierKey *precomputes.PreparedAuthVerifyingKey[fr_bls12381.Vector, *curve_bls12381.G1Jac, *groth16_bls12381.Proof]
	Prepared    *curve_bls12381.G1Jac
}

func DefaultVerifier(gpk *zkbanw.GroupPublicKey, vk *groth16.VerifyingKey) (*Verifier, error) {
	period := utils.Period()

	verifier := Verifier{
		GroupPublicKey: gpk,
		VerifyingKey:   vk,
	}

	_, err := verifier.Update(period)
	if err != nil {
		return nil, err
	}

	return &verifier, nil

}
