package main

import (
	"fmt"
	"strconv"

	corecore "github.com/akakou/zk-ban-system/core"
	serv "github.com/akakou/zk-ban-system/serv/gm"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban-system/utils/codec"
	"github.com/labstack/echo/v4"
)

func RevokeEndpoint[T any](gmServ *serv.GMServer[T]) func(c echo.Context) error {
	return func(c echo.Context) error {
		revoked := c.QueryParam("signature")
		signPeriod := c.QueryParam("period")

		var signature corecore.Signature
		err := codec.UnmarshalBase64(revoked, &signature)
		if err != nil {
			return err
		}

		p, err := strconv.Atoi(signPeriod)
		if err != nil {
			return err
		}

		revokePeriod := utils.Period()
		fmt.Printf("revoked: sign period is %v, and revoked period is %v", signPeriod, revokePeriod)

		gmServ.GM.DB.Client.Revocation.Create().
			SetNym(signature.Signature.Commit.Nym.Bytes()).
			SetRevokedPeriod(int(revokePeriod)).
			SetSignedPeriod(p).
			SaveX(*gmServ.GM.DB.Ctx)

		return c.String(200, "ok")
	}
}
