package gm

import (
	"encoding/base64"
	"errors"

	corecore "github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/core/gm/ent/credential"
	"github.com/akakou/zk-ban-system/utils"
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
	period := utils.Period()

	err = req.JoinRequest.Verify(period, gm.JoinVerifyKey.VerifyingKey)
	if err != nil {
		return "", err
	}

	cred, err := gm.GroupSecretKey.IssueCredential(req.JoinRequest.UserPublicKey)
	if err != nil {
		return "", err
	}

	gm.DB.Client.Credential.Create().
		SetCredential(cred.Signature).
		SetPublicKey(req.JoinRequest.UserPublicKey.Number.Bytes()).
		SetIdentifier(identifer).
		SaveX(*gm.DB.Ctx)

	resp := base64.URLEncoding.EncodeToString(cred.Signature)
	return resp, nil
}
