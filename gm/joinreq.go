package gm

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"

	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/gm/ent/credential"
	"github.com/labstack/echo/v4"
)

var ErrAlreadyRegisterd = errors.New("already account registered")

func IssueCredential[T any](gm *GroupManager[T]) func(c echo.Context) error {
	return func(c echo.Context) error {
		reqBody, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		var req core.JoinRequest[T]
		err = json.Unmarshal(reqBody, &req)
		if err != nil {
			return err
		}

		identifer, err := gm.AuthToken(&req)
		if err != nil {
			return err
		}

		exist, err := gm.DB.Client.Credential.
			Query().
			Where(credential.Identifier(identifer)).
			Exist(*gm.DB.Ctx)

		if err != nil {
			return err
		}

		if exist {
			return ErrAlreadyRegisterd
		}

		cred, err := core.IssueCredential(&req, gm.GroupSecretKey, gm.JoinVerifyKey)
		if err != nil {
			return err
		}

		gm.DB.Client.Credential.Create().
			SetCredential(cred).
			SetPublicKey(req.UserPublicKey).
			SetIdentifier(identifer).
			SaveX(*gm.DB.Ctx)

		resp := base64.URLEncoding.EncodeToString(cred)
		return c.String(200, resp)
	}
}
