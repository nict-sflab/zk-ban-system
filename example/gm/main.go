package main

import (
	"strconv"

	corecore "github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/core/gm"
	core "github.com/akakou/zk-ban-system/core/gm"
	serv "github.com/akakou/zk-ban-system/serv/gm"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/witness"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	gsk, _, err := witness.RandomGroupKeyPair()
	if err != nil {
		panic(err)
	}

	g, err := gm.Default[string](gsk.Bytes(), &core.DBConfig{
		Type:   "sqlite3",
		Config: "file::memory:?cache=shared&_fk=1",
	})
	if err != nil {
		panic(err)
	}

	gmServ := serv.GMServer[string]{
		GM: g,
		AuthToken: func(t *corecore.JoinRequest[string]) (string, error) {
			return "token", nil
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

	gmServ.SetupEchoServer(e)
	e.Debug = true
	e.Use(middleware.Logger())
	e.Start(":8080")
}
