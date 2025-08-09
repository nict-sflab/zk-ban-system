package gm

import "github.com/labstack/echo/v4"

const ISSUE_CREDENTIAL_PATH = "/issue-credential"
const UPDATE_CREDENTIAL_PATH = "/update-credential"
const GROUP_PUBLIC_KEY_PATH = "/group-public-key"
const REVOCATION_LIST_PATH = "/revocation-list"

func (serv *GMServer[T]) SetupEchoServer(e *echo.Echo) {
	issueCred := serv.IssueCredentialEndpoint()
	updateCred := serv.UpdateCredentialEndpoint()
	groupPublicKey := serv.GroupPublicKeyEndpoint()
	revocationList := serv.RevocationList()

	e.POST(ISSUE_CREDENTIAL_PATH, issueCred)
	e.POST(UPDATE_CREDENTIAL_PATH, updateCred)
	e.GET(GROUP_PUBLIC_KEY_PATH, groupPublicKey)
	e.GET(REVOCATION_LIST_PATH, revocationList)
}
