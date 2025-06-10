package verifier

import (
	corecore "github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban/highlevel"
)

func (verifier *Verifier) Verify(signature *corecore.Signature) error {
	err := highlevel.Verify(signature.Signature, signature.Message, verifier.GroupPublicKey, verifier.SignVerifyKey)
	if err != nil {
		return err
	}

	verifier.DB.Client.Pseudonyms.Create().
		SetCount(int(signature.Signature.Counter)).
		SetNym(signature.Signature.Nym).
		SetPeriod(int(signature.Signature.Period))

	return nil
}
