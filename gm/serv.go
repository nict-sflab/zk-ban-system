package gm

import "github.com/labstack/echo/v4"

func SetupEchoServer(e *echo.Echo, g *GroupManager) {
	issueCred := IssueCredential(g)
	updateCred := UpdateCredential(g)

	e.POST("/issue-credential", issueCred)
	e.POST("/update-credential", updateCred)
}
