package gm

import (
	"encoding/base64"
	"errors"

	corecore "github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/core/gm/ent/credential"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/highlevel"
	"github.com/akakou/zk-ban/witness"
)

var ErrAlreadyIssueCredential = errors.New("already issue credential")

func (gm *GroupManager[T]) UpdateCredential(req *corecore.UpdateRequest) (string, error) {
	rl := witness.RevocationList{}

	exist, err := gm.DB.Client.Credential.
		Query().
		Where(credential.PublicKey(req.UserPublicKey)).
		Exist(*gm.DB.Ctx)

	if err != nil {
		return "", err
	}

	if exist {
		return "", ErrAlreadyIssueCredential
	}

	after := utils.Today()

	index := DoubleMapKey{
		First:  int(after),
		Second: int(req.Before),
	}
	verifier, hasVerifier := PreparedSnarkVerifiers[index]

	if !hasVerifier {
		prepared, err := highlevel.PrepareVerification(rl, gm.UpdateVerifyKey)
		if err != nil {
			return "", err
		}

		verifier = &PreparedSnarkVerifier{
			VerifierKey: PreparableSnarkVerifierKey{
				VerifyingKey: gm.UpdateVerifyKey,
				Size: &RevocationListSize{
					0, 0,
				},
			},
			Prepared: prepared,
		}

		PreparedSnarkVerifiers[index] = verifier
	}

	err = highlevel.VerifyUpdateRequest(req.Proof, req.UserPublicKey, after, req.Before, rl, gm.GroupPublicKey, verifier.Prepared, gm.UpdateVerifyKey)
	if err != nil {
		return "", err
	}

	cred, err := highlevel.IssueCredential(after, req.UserPublicKey, gm.GroupSecretKey)
	if err != nil {
		return "", err
	}

	gm.DB.Client.Credential.Create().
		SetCredential(cred).
		SetPublicKey(req.UserPublicKey).
		SaveX(*gm.DB.Ctx)

	resp := base64.URLEncoding.EncodeToString(cred)
	return resp, nil
}
