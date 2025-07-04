package verifier

import (
	corecore "github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/load"
	"github.com/akakou/zk-ban/highlevel"
)

func (verifier *Verifier) Verify(signature *corecore.Signature) error {
	err := highlevel.Verify(signature.Signature, signature.Message, verifier.GroupPublicKey, load.SignVerifierKey)
	if err != nil {
		return err
	}

	return nil
}
