package gm

import (
	"encoding/json"
	"io"

	"github.com/akakou/zk-ban-system/core/core"
	"github.com/labstack/echo/v4"
)

func (serv *GMServer[T]) IssueCredentialEndpoint() func(c echo.Context) error {
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

		identifer, err := serv.AuthToken(&req)
		if err != nil {
			return err
		}

		resp, err := serv.GM.IssueCredential(identifer, &req)
		if err != nil {
			return err
		}

		return c.String(200, resp)
	}
}
