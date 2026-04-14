package verifier

import (
	"github.com/akakou/zk-ban/precomputes"
)

func (verifier *Verifier) Update(period int64) (*PreparedAuthSnarkVerifier, error) {
	vk, err := precomputes.NewAuthVerificationKeyBLS12381(*verifier.VerifyingKey)
	if err != nil {
		return nil, err
	}

	prepared, err := vk.PrecomputeVerify(period, verifier.GroupPublicKey)
	if err != nil {
		return nil, err
	}

	preparedVK := PreparedAuthSnarkVerifier{
		VerifierKey: vk,
		Prepared:    *prepared,
	}

	verifier.PreparedVerifyingKey = &preparedVK
	verifier.Period = period

	return &preparedVK, nil
}
