package gm

import (
	"encoding/base64"
	"errors"

	corecore "github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/core/gm/ent/idtoken"
)

var ErrAlreadyRegisterd = errors.New("already account registered")

func (gm *GroupManager[T]) IssueCredential(identifer string, req *corecore.JoinRequest[T], period int64) (string, error) {
	exist, err := gm.DB.Client.IDToken.
		Query().
		Where(idtoken.IdTokenEQ(identifer)).
		Exist(*gm.DB.Ctx)

	if err != nil {
		return "", err
	}

	if exist {
		return "", ErrAlreadyRegisterd
	}

	err = req.JoinRequest.Verify(period, *gm.JoinVerifyKey)
	if err != nil {
		return "", err
	}

	cred, err := gm.GroupSecretKey.IssueCredential(req.JoinRequest.UserPublicKey)
	if err != nil {
		return "", err
	}

	gm.DB.Client.IDToken.Create().
		SetIdToken(identifer).
		SaveX(*gm.DB.Ctx)

	gm.DB.Client.Credential.Create().
		SetCredential(cred.Signature).
		SetPublicKey(req.JoinRequest.UserPublicKey.Number.Bytes()).
		SaveX(*gm.DB.Ctx)

	resp := base64.URLEncoding.EncodeToString(cred.Signature)
	return resp, nil
}
