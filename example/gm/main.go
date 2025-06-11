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

const PATH = "../../../zk-ban-android/tools/join_verifier.json"

func main() {
	gmDB, err := core.NewDB(&core.DBConfig{
		Type:   "sqlite3",
		Config: "file::memory:?cache=shared&_fk=1",
	})

	if err != nil {
		panic(err)
	}

	dat, err := os.ReadFile(PATH)
	if err != nil {
		panic(err)
	}

	gsk, gpk, err := witness.RandomGroupKeyPair()
	if err != nil {
		panic(err)
	}

	g := core.GroupManager[string]{
		GroupSecretKey: gsk.Bytes(),
		GroupPublicKey: gpk.Bytes(),
		JoinVerifyKey:  dat,
		DB:             gmDB,
	}

	gmServ := serv.GMServer[string]{
		GM: &g,
		AuthToken: func(t *corecore.JoinRequest[string]) (string, error) {
			return "token", nil
		},
	}

	e := echo.New()
	gmServ.SetupEchoServer(e)
	e.Debug = true
	e.Use(middleware.Logger())
	e.Start(":8080")
}
