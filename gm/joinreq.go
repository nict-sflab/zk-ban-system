package gm

import (
	"encoding/base64"
	"encoding/json"
	"io"

	"github.com/akakou/zk-ban-system/core"
	"github.com/labstack/echo/v4"
)

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

		err = gm.CheckToken(&req)
		if err != nil {
			return err
		}

		cred, err := core.IssueCredential(&req, gm.GroupSecretKey, gm.JoinVerifyKey)
		if err != nil {
			return err
		}

		resp := base64.URLEncoding.EncodeToString(cred)
		return c.String(200, resp)
	}
}
