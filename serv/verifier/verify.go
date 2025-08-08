package verifier

import (
	"encoding/json"

	"github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/labstack/echo/v4"
)

func (serv *VerifierServer) VerifyEndpoint() func(c echo.Context) error {
	return func(c echo.Context) error {
		sigString := c.QueryParam("signature")
		period := utils.Today()

		var signature core.Signature
		err := json.Unmarshal([]byte(sigString), &signature)
		if err != nil {
			return err
		}

		// nym := base64.URLEncoding.EncodeToString(signature.Signature.Commit.Nym)
		// fmt.Printf("%v\n%v\n", nym, signature.Count)

		err = serv.Verifier.Verify(&signature, period)
		if err != nil {
			return err
		}

		return c.String(200, "ok")
	}
}
