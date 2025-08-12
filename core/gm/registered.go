package gm

import (
	"encoding/base64"

	"github.com/akakou/zk-ban-system/core/gm/ent/credential"
)

func (gm *GroupManager[T]) queryAlreadyRegisteredCredential(upk []byte) (string, error) {
	existCred, err := gm.DB.Client.Credential.Query().
		Where(
			credential.PublicKey(upk),
		).Only(*gm.DB.Ctx)

	if err == nil {
		resp := base64.URLEncoding.EncodeToString(existCred.Credential)
		return resp, nil
	} else if err.Error() != ErrCredentialNotFound.Error() {
		return "", err
	}

	return "", nil
}
