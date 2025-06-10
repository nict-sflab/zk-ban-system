package verifier

import (
	"errors"

	corecore "github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/core/verifier/ent/pseudonyms"
	"github.com/akakou/zk-ban/highlevel"
)

var errSameNymExist = errors.New("same pesudonym exist")

func (verifier *Verifier) Verify(signature *corecore.Signature) error {
	hasExist := verifier.DB.Client.Pseudonyms.Query().
		Where(pseudonyms.NymEQ(signature.Signature.Nym)).
		ExistX(*verifier.DB.Ctx)

	if hasExist {
		return errSameNymExist
	}

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
