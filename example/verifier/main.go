package main

import (
	"time"

	signer "github.com/akakou/zk-ban-system/client/signer"
	core "github.com/akakou/zk-ban-system/core/verifier"
	"github.com/akakou/zk-ban-system/load"
	serv "github.com/akakou/zk-ban-system/serv/verifier"
	"github.com/akakou/zk-ban-system/utils"
	zkbanw "github.com/akakou/zk-ban/witness"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	utils.PeriodUnit = time.Duration(time.Minute / 2)

	gpkBuf, err := signer.FetchGroupPublicKey("http://localhost:8080/group-public-key")
	if err != nil {
		panic(err)
	}

	gpk, err := zkbanw.GroupPublicKeyFromBytes(gpkBuf)
	if err != nil {
		panic(err)
	}

	verifierKey, err := load.DocodeVerifyingKey(load.SignVerifierKey)
	if err != nil {
		panic(err)
	}

	verifierServ := serv.VerifierServer{
		Verifier: &core.Verifier{
			GroupPublicKey: gpk,
			VerifyingKey:   verifierKey,
		},
	}

	e := echo.New()

	// e.GET("/period", func(c echo.Context) error {
	// 	str := c.QueryParam("period")
	// 	if str == "" {
	// 		return c.String(200, strconv.Itoa(int(utils.Period())))
	// 	}

	// 	i, err := strconv.Atoi(str)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	utils.Period = func() int64 {
	// 		return int64(i)
	// 	}

	// 	return c.String(200, strconv.Itoa(i))
	// })

	verifierServ.SetupEchoServer(e)
	e.Debug = true
	e.Use(middleware.Logger())
	e.Start(":8000")
}
