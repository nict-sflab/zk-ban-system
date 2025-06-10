package verifier

import (
	"encoding/json"
	"io"

	"github.com/akakou/zk-ban-system/core"
	"github.com/labstack/echo/v4"
)

func Verify(verifier *Verifier) func(c echo.Context) error {
	return func(c echo.Context) error {
		reqBody, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		var signature core.Signature

		err = json.Unmarshal(reqBody, &signature)
		if err != nil {
			return err
		}

		err = core.Verify(&signature, verifier.GroupPublicKey, verifier.SignVerifyKey)

		if err != nil {
			return err
		}

		verifier.DB.Client.Pseudonyms.Create().
			SetCount(int(signature.Signature.Counter)).
			SetNym(signature.Signature.Nym).
			SetPeriod(int(signature.Signature.Period))

		return nil
	}
}
