package gm

import (
	"encoding/base64"

	"github.com/akakou/zk-ban/witness"
	"github.com/labstack/echo/v4"
)

func (serv *GMServer[T]) GroupPublicKeyEndpoint() func(c echo.Context) error {
	return func(c echo.Context) error {
		gpk := base64.URLEncoding.EncodeToString(serv.GM.GroupPublicKey)
		return c.String(200, gpk)
	}
}

func (serv *GMServer[T]) RevocationList() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(200, witness.RevocationList{})
	}
}
