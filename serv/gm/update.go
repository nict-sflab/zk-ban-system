package gm

import (
	"io"

	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban-system/utils/codec"
	"github.com/labstack/echo/v4"
)

func (serv GMServer[T]) UpdateCredentialEndpoint() func(c echo.Context) error {
	return func(c echo.Context) error {
		reqBody, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		var req core.UpdateRequest
		err = codec.Unmarshal(reqBody, &req)
		if err != nil {
			return err
		}

		resp, err := serv.GM.UpdateCredential(&req, utils.Period())
		if err != nil {
			return err
		}

		return c.String(200, resp)
	}
}
