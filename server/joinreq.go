package server

import (
	"encoding/base64"
	"io"

	"github.com/akakou/zk-ban-system/core"
	"github.com/labstack/echo/v4"
)

type GroupManager struct {
	GroupSecretKey  []byte
	GroupPublicKey  []byte
	JoinVerifyKey   []byte
	UpdateVerifyKey []byte
}

func IssueCredential(gm *GroupManager) func(c echo.Context) error {
	return func(c echo.Context) error {
		reqBody, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		cred, err := core.IssueCredential[string](reqBody, gm.GroupSecretKey, gm.JoinVerifyKey)
		if err != nil {
			return err
		}

		resp := base64.URLEncoding.EncodeToString(cred)
		return c.String(200, resp)
	}
}
