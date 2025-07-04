package gm

import (
	"encoding/base64"
	"errors"

	corecore "github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/core/gm/ent/credential"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/highlevel"
)

var ErrAlreadyIssueCredential = errors.New("already issue credential")

func (gm *GroupManager[T]) UpdateCredential(req *corecore.UpdateRequest) (string, error) {
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

	index := corecore.KeyIndex{
		First:  after,
		Second: req.Before,
	}

	verifier, hasVerifier := gm.PreparedSnarkVerifiers[index]

	if !hasVerifier {
		rl, err := gm.QueryRL(req.Before)
		if err != nil {
			return "", err
		}

		key := gm.VerifierKeys[*rl.Size]
		prepared, err := highlevel.PrepareVerification(rl.List, *key)
		if err != nil {
			return "", err
		}

		verifier = &PreparedSnarkVerifier{
			VerifierKey: key,
			Prepared:    prepared,
		}

		gm.PreparedSnarkVerifiers[index] = verifier
	}

	err = highlevel.VerifyUpdateRequest(req.Proof, req.UserPublicKey, after, req.Before, gm.GroupPublicKey, verifier.Prepared, *verifier.VerifierKey)
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
