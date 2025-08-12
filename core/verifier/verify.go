package verifier

import (
	"fmt"

	corecore "github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban/primitives"
)

func (verifier *Verifier) Verify(signature *corecore.Signature, period int64) error {
	if signature.Count <= 0 || signature.Count > verifier.CountMax {
		return fmt.Errorf("count does not fulfill the size requirement: 0 < %d(count) <= %d (max)", signature.Count, verifier.CountMax)
	}
	err := signature.Signature.Verify(
		primitives.BigIntFromBytes(signature.Message),
		signature.Count,
		period,
		verifier.GroupPublicKey,
		verifier.VerifyingKey.VerifyingKey)
	if err != nil {
		return err
	}

	return nil
}
