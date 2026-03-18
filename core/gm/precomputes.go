package gm

import (
	corecore "github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban/precomputes"
)

func (gm *GroupManager[T]) readyUpdateVerifyKey(index corecore.KeyIndex) (*PreparedSnarkVerifier, error) {
	before := index.Second
	after := index.First

	rl, v, err := gm.QueryRL(before, after)
	if err != nil {
		return nil, err
	}

	vk, err := precomputes.NewUpdateVerificationKeyBLS12381(*v.VerifyKey)
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

	gm.PreparedSnarkVerifiers[index] = &verifier

	return &verifier, nil
}

func (gm *GroupManager[T]) readyUpdateVerifyKeysForRange(periodRange, after int64) error {
	before := after - periodRange

	for i := before; i < after; i++ {
		_, err := gm.readyUpdateVerifyKey(corecore.KeyIndex{
			First:  after,
			Second: before,
		})

		if err != nil {
			return err
		}
	}

	return nil
}
