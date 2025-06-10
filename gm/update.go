package gm

import (
	"encoding/base64"
	"io"

	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban/witness"
	"github.com/labstack/echo/v4"
)

func UpdateCredential(gm *GroupManager) func(c echo.Context) error {
	return func(c echo.Context) error {
		rl := witness.RevocationList{}

		reqBody, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		cred, err := core.UpdateCredential(reqBody, gm.GroupSecretKey, gm.GroupPublicKey, rl, gm.UpdateVerifyKey)
		if err != nil {
			return err
		}

		resp := base64.URLEncoding.EncodeToString(cred)
		return c.String(200, resp)
	}
}
