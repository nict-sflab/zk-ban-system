package gm

import "github.com/labstack/echo/v4"

func (serv *GMServer[T]) SetupEchoServer(e *echo.Echo) {
	issueCred := serv.IssueCredentialEndpoint()
	updateCred := serv.UpdateCredentialEndpoint()
	groupPublicKey := serv.GroupPublicKeyEndpoint()

	e.POST("/issue-credential", issueCred)
	e.POST("/update-credential", updateCred)
	e.GET("/group-public-key", groupPublicKey)
}
