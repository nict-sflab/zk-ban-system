package verifier

import (
	"fmt"

	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban-system/utils/codec"
	"github.com/labstack/echo/v4"
)

func (serv *VerifierServer) VerifyEndpoint() func(c echo.Context) error {
	return func(c echo.Context) error {
		sigString := c.QueryParam("signature")
		period := utils.Period()

		var signature core.Signature
		err := codec.UnmarshalBase64(sigString, &signature)
		if err != nil {
			return err
		}

		fmt.Printf("%v\nPeriod: %v\nCount: %v\n", sigString, period, signature.Count)

		err = serv.Verifier.Verify(&signature, period)
		if err != nil {
			return err
		}

		return c.String(200, "ok")
	}
}
