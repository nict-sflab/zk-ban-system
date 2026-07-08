package main

import (
	"fmt"
	"log"

	"github.com/akakou/zk-ban-system/core/gm"
	core "github.com/akakou/zk-ban-system/core/gm"
	"github.com/akakou/zk-ban-system/example/gm/extension"
	serv "github.com/akakou/zk-ban-system/serv/gm"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/dump"
	"github.com/akakou/zk-ban/witness"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	dump.KeyPath = "../../dump/"
	utils.PeriodUnit = 120000000000
	serv.PeriodRange = 10
	fmt.Printf("unit: %v\n", int64(utils.PeriodUnit))

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

	e.GET("/revoke", RevokeEndpoint(&gmServ))

	// extension for test
	// curl http://localhost:8080/bulk_revoke?T=30&deltaL=30000&shape_type=0
	e.GET("/bulk_revoke", extension.BulkRevokeEndpoint(&gmServ))

	e.Static("/admin", "./static")

	gmServ.SetupEchoServer(e)
	e.Debug = true
	e.Use(middleware.Logger())

	go gmServ.RunKeyPrecomputeDaemon()
	panic(e.Start(":8080"))
}
