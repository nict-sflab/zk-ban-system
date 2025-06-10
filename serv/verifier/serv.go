package verifier

import "github.com/labstack/echo/v4"

func (serv *VerifierServer) SetupEchoServer(e *echo.Echo) {
	endpoint := serv.VerifyEndpoint()

	e.POST("/verify", endpoint)
}
