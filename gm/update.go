package gm

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"

	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/gm/ent/credential"
	"github.com/akakou/zk-ban/witness"
	"github.com/labstack/echo/v4"
)

var ErrAlreadyIssueCredential = errors.New("already issue credential")

func UpdateCredential[T any](gm *GroupManager[T]) func(c echo.Context) error {
	return func(c echo.Context) error {
		rl := witness.RevocationList{}

		reqBody, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		var req core.UpdateRequest
		err = json.Unmarshal(reqBody, &req)
		if err != nil {
			return err
		}

		exist, err := gm.DB.Client.Credential.
			Query().
			Where(credential.PublicKey(req.UserPublicKey)).
			Exist(*gm.DB.Ctx)

		if err != nil {
			return err
		}

		if exist {
			return ErrAlreadyIssueCredential
		}

		cred, err := core.UpdateCredential(&req, gm.GroupSecretKey, gm.GroupPublicKey, rl, gm.UpdateVerifyKey)
		if err != nil {
			return err
		}

		gm.DB.Client.Credential.Create().
			SetCredential(cred).
			SetPublicKey(req.UserPublicKey).
			SaveX(*gm.DB.Ctx)

		resp := base64.URLEncoding.EncodeToString(cred)
		return c.String(200, resp)
	}
}
