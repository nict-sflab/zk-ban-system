package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	corecore "github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/core/gm"
	core "github.com/akakou/zk-ban-system/core/gm"
	serv "github.com/akakou/zk-ban-system/serv/gm"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/dump"
	"github.com/akakou/zk-ban/witness"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	dump.KeyPath = "../../dump/"
	utils.PeriodUnit = utils.HalfMinutes
	gsk, _, err := witness.RandomGroupKeyPair()
	if err != nil {
		log.Fatal(err)
	}

	g, err := gm.Default[string](gsk.Bytes(), &core.DBConfig{
		Type:   "sqlite3",
		Config: "file::memory:?cache=shared&_fk=1",
	})
	if err != nil {
		log.Fatal(err)
	}

	gmServ := serv.GMServer[string]{
		GM:        g,
		AuthToken: authToken(),
	}

	e := echo.New()

	e.GET("/revoke", func(c echo.Context) error {
		revoked := c.QueryParam("signature")
		signPeriod := c.QueryParam("period")

		var signature corecore.Signature
		err := json.Unmarshal([]byte(revoked), &signature)
		if err != nil {
			return err
		}

		p, err := strconv.Atoi(signPeriod)
		if err != nil {
			return err
		}

		revokePeriod := utils.Period()
		fmt.Printf("revoked: sign period is %v, and revoked period is %v", signPeriod, revokePeriod)

		gmServ.GM.DB.Client.Revocation.Create().
			SetCount(int(signature.Count)).
			SetNym(signature.Signature.Commit.Nym.Bytes()).
			SetRevokedPeriod(int(revokePeriod)).
			SetSignedPeriod(p).
			SaveX(*gmServ.GM.DB.Ctx)

		return c.String(200, "ok")
	})

	e.Static("/admin", "./static")

	gmServ.SetupEchoServer(e)
	e.Debug = true
	e.Use(middleware.Logger())
	panic(e.Start(":8080"))
}
