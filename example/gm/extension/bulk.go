package extension

import (
	"fmt"
	"strconv"

	serv "github.com/akakou/zk-ban-system/serv/gm"

	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/test"
	"github.com/akakou/zk-ban/witness"
	"github.com/labstack/echo/v4"
)

func BulkRevokeEndpoint[T any](gmServ *serv.GMServer[T]) func(c echo.Context) error {
	return func(c echo.Context) error {
		paramT := c.QueryParam("T")
		paramDeltaL := c.QueryParam("deltaL")
		paramShapeType := c.QueryParam("shape")

		fmt.Printf("bulk revoke: \nT=%v\ndeltaL=%v\n\n'", paramT, paramDeltaL)
		t, err := strconv.Atoi(paramT)
		if err != nil {
			return err
		}

		deltaL, err := strconv.Atoi(paramDeltaL)
		if err != nil {
			return err
		}

		shapeType, err := strconv.Atoi(paramShapeType)
		if err != nil {
			shapeType = 0
		}

		rlU := test.EmptyUniformRevocationList(t, deltaL)
		rlP := test.EmptyProportionalRevocationList(t, deltaL)
		rlG := test.EmptyGaussianRevocationList(t, deltaL)
		witness.InitBigInt = witness.ZeroInitBigInt

		rl := []witness.RevocationList{rlU, rlP, rlG}[shapeType]

		revokePeriod := int(utils.Period())

		for i, r := range rl {
			signedPeriod := revokePeriod - t + i - 1
			for _j, nym := range r.Nyms {
				gmServ.GM.DB.Client.Revocation.Create().
					SetNym(nym.Bytes()).
					SetRevokedPeriod(revokePeriod).
					SetSignedPeriod(signedPeriod).
					SaveX(*gmServ.GM.DB.Ctx)

				fmt.Printf("%v, %v\n", i, _j)
			}
		}
		return c.String(200, "ok")
	}
}
