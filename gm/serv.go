package gm

import "github.com/labstack/echo/v4"

func (g *GroupManager[T]) SetupEchoServer(e *echo.Echo) {
	issueCred := IssueCredential(g)
	updateCred := UpdateCredential(g)

	e.POST("/issue-credential", issueCred)
	e.POST("/update-credential", updateCred)
}
