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

	period := utils.Today()

	err = highlevel.VerifyUpdateRequest(req.Proof, req.UserPublicKey, period, req.Before, rl, gm.GroupPublicKey, gm.UpdateVerifyKey)
	if err != nil {
		return "", err
	}

	cred, err := highlevel.IssueCredential(period, req.UserPublicKey, gm.GroupSecretKey)
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
