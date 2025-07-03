package verifier

import (
	corecore "github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/dump"
	"github.com/akakou/zk-ban/highlevel"
)

func (verifier *Verifier) Verify(signature *corecore.Signature) error {
	err := highlevel.Verify(signature.Signature, signature.Message, verifier.GroupPublicKey, dump.SignVerifierKey)
	if err != nil {
		return err
	}

	return nil
}
