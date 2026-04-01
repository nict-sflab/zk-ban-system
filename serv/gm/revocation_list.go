package gm

import (
	"strconv"

	"github.com/akakou/zk-ban-system/utils"
	"github.com/labstack/echo/v4"
)

func (serv *GMServer[T]) RevocationList() func(c echo.Context) error {
	return func(c echo.Context) error {
		beforeQueryParam := c.QueryParam("before")
		before, err := strconv.Atoi(beforeQueryParam)
		if err != nil {
			return err
		}

		after := utils.Period()

		rl, err := serv.GM.RevocationList(int64(before), int64(after))
		if err != nil {
			return err
		}

		return c.Blob(200, "application/octet-stream", rl)
	}
}
