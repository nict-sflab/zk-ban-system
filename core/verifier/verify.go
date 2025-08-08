package verifier

import (
	corecore "github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban/primitives"
)

func (verifier *Verifier) Verify(signature *corecore.Signature, period int64) error {
	err := signature.Signature.Verify(
		primitives.BigIntFromBytes(signature.Message),
		signature.Count,
		period,
		verifier.GroupPublicKey,
		verifier.VerifyingKey)
	if err != nil {
		return err
	}

	return nil
}
