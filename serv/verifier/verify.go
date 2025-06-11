package verifier

import (
	"encoding/json"

	"github.com/akakou/zk-ban-system/core/core"
	"github.com/labstack/echo/v4"
)

func (serv *VerifierServer) VerifyEndpoint() func(c echo.Context) error {
	return func(c echo.Context) error {
		sigString := c.QueryParam("signature")

		var signature core.Signature
		err := json.Unmarshal([]byte(sigString), &signature)
		if err != nil {
			return err
		}

		err = serv.Verifier.Verify(&signature)
		if err != nil {
			return err
		}

		return c.String(200, "ok")
	}
}
