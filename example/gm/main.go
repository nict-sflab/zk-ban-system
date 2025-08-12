package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	corecore "github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/core/gm"
	core "github.com/akakou/zk-ban-system/core/gm"
	serv "github.com/akakou/zk-ban-system/serv/gm"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/witness"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	utils.PeriodUnit = time.Duration(time.Minute / 2)
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

	ctx := context.Background()
	svc, err := NewFirebaseAuthService(ctx, "serviceAccount.json")
	if err != nil {
		log.Fatal(err)
	}

	gmServ := serv.GMServer[string]{
		GM:        g,
		AuthToken: svc.FirebaseAuth(),
		// AuthToken: allOKAuth,
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
	e.Start(":8080")
}
