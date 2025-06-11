package main

import (
	"os"
	"strconv"

	signer "github.com/akakou/zk-ban-system/client/signer"
	core "github.com/akakou/zk-ban-system/core/verifier"
	serv "github.com/akakou/zk-ban-system/serv/verifier"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const SIGN_VERIFIER_PATH = "../../../zk-ban-android/tools/sign_verifier.json"

func main() {
	db, err := core.NewDB(&core.DBConfig{
		Type:   "sqlite3",
		Config: "file::memory:?cache=shared&_fk=1",
	})

	if err != nil {
		panic(err)
	}

	signVerifier, err := os.ReadFile(SIGN_VERIFIER_PATH)
	if err != nil {
		panic(err)
	}

	gpk, err := signer.FetchGroupPublicKey("http://localhost:8080/group-public-key")
	if err != nil {
		panic(err)
	}

	verifierServ := serv.VerifierServer{
		Verifier: &core.Verifier{
			SignVerifyKey:  signVerifier,
			GroupPublicKey: gpk,
			DB:             db,
		},
	}

	e := echo.New()

	e.GET("/period", func(c echo.Context) error {
		str := c.QueryParam("period")
		i, err := strconv.Atoi(str)
		if err != nil {
			return err
		}

		utils.Today = func() int64 {
			return int64(i)
		}

		return c.String(200, strconv.Itoa(i))
	})

	verifierServ.SetupEchoServer(e)
	e.Debug = true
	e.Use(middleware.Logger())
	e.Start(":8000")
}
