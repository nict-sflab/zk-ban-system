package gm

import (
	"encoding/base64"

	"github.com/labstack/echo/v4"
)

func (serv *GMServer[T]) GroupPublicKeyEndpoint() func(c echo.Context) error {
	return func(c echo.Context) error {
		gpk := base64.URLEncoding.EncodeToString(serv.GM.GroupPublicKey.Bytes())
		return c.String(200, gpk)
	}
}
