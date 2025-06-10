package verifier

import (
	"encoding/json"
	"io"

	"github.com/akakou/zk-ban-system/core/core"
	"github.com/labstack/echo/v4"
)

func VerifyEndpoint(serv *VerifierServer) func(c echo.Context) error {
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

		return nil
	}
}
