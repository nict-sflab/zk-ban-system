package main

import (
	"os"

	corecore "github.com/akakou/zk-ban-system/core/core"
	core "github.com/akakou/zk-ban-system/core/gm"
	serv "github.com/akakou/zk-ban-system/serv/gm"
	"github.com/akakou/zk-ban/witness"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const JOIN_VERIFIER_PATH = "../../../zk-ban-android/tools/join_verifier.json"
const UPDATE_VERIFIER_PATH = "../../../zk-ban-android/tools/update_verifier.json"

func main() {
	gmDB, err := core.NewDB(&core.DBConfig{
		Type:   "sqlite3",
		Config: "file::memory:?cache=shared&_fk=1",
	})

	if err != nil {
		panic(err)
	}

	joinVerifier, err := os.ReadFile(JOIN_VERIFIER_PATH)
	if err != nil {
		panic(err)
	}

	updateVerifier, err := os.ReadFile(UPDATE_VERIFIER_PATH)
	if err != nil {
		panic(err)
	}

	gsk, gpk, err := witness.RandomGroupKeyPair()
	if err != nil {
		panic(err)
	}

	g := core.GroupManager[string]{
		GroupSecretKey:  gsk.Bytes(),
		GroupPublicKey:  gpk.Bytes(),
		JoinVerifyKey:   joinVerifier,
		UpdateVerifyKey: updateVerifier,
		DB:              gmDB,
	}

	gmServ := serv.GMServer[string]{
		GM: &g,
		AuthToken: func(t *corecore.JoinRequest[string]) (string, error) {
			return "token", nil
		},
	}

	e := echo.New()
	e.GET("/reset", func(c echo.Context) error {
		gmServ.GM.DB.Close()

		gmDB, err := core.NewDB(&core.DBConfig{
			Type:   "sqlite3",
			Config: "file::memory:?cache=shared&_fk=1",
		})
		if err != nil {
			panic(err)
		}

		gmServ.GM.DB = gmDB

		return c.String(200, "reset")
	})

	gmServ.SetupEchoServer(e)
	e.Debug = true
	e.Use(middleware.Logger())
	e.Start(":8080")
}
