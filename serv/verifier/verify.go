package verifier

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

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

		nym := base64.URLEncoding.EncodeToString(signature.Signature.Nym)
		fmt.Printf("%v\n%v\n%v", nym, signature.Signature.Counter, signature.Signature.Period)

		err = serv.Verifier.Verify(&signature)
		if err != nil {
			return err
		}

		return c.String(200, "ok")
	}
}
