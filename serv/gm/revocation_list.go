package gm

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func (serv *GMServer[T]) RevocationList() func(c echo.Context) error {
	return func(c echo.Context) error {
		beforeQueryParam := c.QueryParam("before")
		before, err := strconv.Atoi(beforeQueryParam)
		if err != nil {
			return err
		}

		rl, err := serv.GM.RevocationList(int64(before))
		if err != nil {
			return err
		}

		return c.String(200, rl)
	}
}
