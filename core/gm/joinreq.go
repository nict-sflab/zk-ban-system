package gm

import (
	"encoding/base64"
	"errors"

	corecore "github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/core/gm/ent/credential"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/highlevel"
)

var ErrAlreadyRegisterd = errors.New("already account registered")

func (gm *GroupManager[T]) IssueCredential(identifer string, req *corecore.JoinRequest[T]) (string, error) {
	exist, err := gm.DB.Client.Credential.
		Query().
		Where(credential.Identifier(identifer)).
		Exist(*gm.DB.Ctx)

	if err != nil {
		return "", err
	}

	if exist {
		return "", ErrAlreadyRegisterd
	}
	period := utils.Today()

	err = highlevel.VerifyJoinReq(req.Proof, req.UserPublicKey, period, gm.JoinVerifyKey)
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
		SetIdentifier(identifer).
		SaveX(*gm.DB.Ctx)

	resp := base64.URLEncoding.EncodeToString(cred)
	return resp, nil
}
