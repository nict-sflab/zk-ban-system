package gm

import "github.com/akakou/zk-ban/precomputes"

func (gm *GroupManager[T]) precomputesVerifyUpdateRequest(before, after int64) (*PreparedSnarkVerifier, error) {
	rl, v, err := gm.QueryRL(before, after)
	if err != nil {
		return nil, err
	}

	vk, err := precomputes.NewUpdateVerificationKeyBLS12381(v.VerifyKey.VerifyingKey)
	if err != nil {
		return nil, err
	}

	prepared, err := vk.PrecomputeVerify(*rl.List, &gm.GroupPublicKey)
	if err != nil {
		return nil, err
	}

	verifier := PreparedSnarkVerifier{
		VerifierKey: vk,
		Prepared:    *prepared,
	}

	return &verifier, nil
}
