package verifier

import (
	corecore "github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban/primitives"
)

func (verifier *Verifier) Verify(signature *corecore.Signature, period int64) error {
	err := verifier.PreparedVerifyingKey.VerifierKey.VerifyPrepared(
		verifier.PreparedVerifyingKey.Prepared,
		primitives.BigIntFromBytes(signature.Message), signature.Signature)

	return err
}
