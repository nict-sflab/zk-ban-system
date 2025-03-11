package server

import (
	"io"

	"github.com/akakou/zk-ban-system/core"
	"github.com/labstack/echo/v4"
)

type GroupManager struct {
	GroupSecretKey []byte
	VerifyKey      []byte
}

func IssueCredential(gm *GroupManager) func(c echo.Context) error {
	return func(c echo.Context) error {
		reqBody, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		cred, err := core.IssueCredential[string](reqBody, gm.GroupSecretKey, gm.VerifyKey)
		if err != nil {
			return err
		}

		return c.JSON(200, cred)
	}
}
