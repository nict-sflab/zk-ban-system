package main

import (
	"log"

	signer "github.com/akakou/zk-ban-system/client/signer"
	core "github.com/akakou/zk-ban-system/core/verifier"
	serv "github.com/akakou/zk-ban-system/serv/verifier"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/load"
	zkbanw "github.com/akakou/zk-ban/witness"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	utils.PeriodUnit = utils.HalfMinutes
	gpkBuf, err := signer.FetchGroupPublicKey("http://localhost:8080/group-public-key")
	if err != nil {
		log.Fatal(err)
	}

	gpk, err := zkbanw.GroupPublicKeyFromBytes(gpkBuf)
	if err != nil {
		log.Fatal(err)
	}

	verifierKey, err := load.LoadBasicGroupManagerKey("update")
	if err != nil {
		log.Fatal(err)
	}

	verifierServ := serv.VerifierServer{
		Verifier: &core.Verifier{
			GroupPublicKey: gpk,
			VerifyingKey:   verifierKey,
			CountMax:       2,
		},
	}

	e := echo.New()

	e.Static("/", "./static")

	verifierServ.SetupEchoServer(e)
	e.Debug = true
	e.Use(middleware.Logger())
	e.Start(":8000")
}
