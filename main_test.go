package zkbansystem

import (
	"bytes"
	"testing"
	"time"

	"github.com/akakou/zk-ban-system/client/signer"
	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/core/gm"
	coreverifier "github.com/akakou/zk-ban-system/core/verifier"
	gmserv "github.com/akakou/zk-ban-system/serv/gm"
	"github.com/akakou/zk-ban-system/serv/verifier"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/dump"
	"github.com/akakou/zk-ban/load"
	"github.com/akakou/zk-ban/witness"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const SLEEP_TIME = 3

var period = 1

func today() int64 {
	return int64(period)
}

func passDay() {
	period += 1
}

func TestMain(t *testing.T) {
	utils.Period = today
	dump.KeyPath = "./dump/"

	gsk, _, err := witness.RandomGroupKeyPair()
	if err != nil {
		t.Fatal(err)
	}

	g, err := gm.Default[string](gsk.Bytes(), &gm.DBConfig{
		Type:   "sqlite3",
		Config: "file::memory:?cache=shared&_fk=1",
	})

	if err != nil {
		t.Fatal(err)
	}

	gmServ := gmserv.GMServer[string]{
		GM: g,
		AuthToken: func(t *core.JoinRequest[string]) (string, error) {
			return t.Option, nil
		},
	}

	e := echo.New()
	e.Use(middleware.Logger())

	gmServ.SetupEchoServer(e)

	verifierKey, err := load.LoadBasicGroupManagerKey("sign")
	if err != nil {
		t.Fatal(err)
	}

	v, err := coreverifier.DefaultVerifier(
		&g.GroupPublicKey,
		verifierKey,
	)
	if err != nil {
		t.Fatal(err)
	}

	verifierServ := verifier.VerifierServer{
		Verifier: v,
	}

	verifierServ.SetupEchoServer(e)

	go func() {
		if err := e.Start(":1323"); err != nil {
			t.Fatal(err)
		}
	}()

	time.Sleep(SLEEP_TIME * time.Second)

	s1, err := signer.RequestJoin("", "http://localhost:1323/issue-credential")
	if err != nil {
		t.Fatal(err)
	}

	s2, err := signer.RequestJoin("2", "http://localhost:1323/issue-credential")
	if err != nil {
		t.Fatal(err)
	}

	ns, err := signer.RequestJoin("", "http://localhost:1323/issue-credential")
	if err == nil {
		t.Fatal(ns, err)
	}

	gpk2, err := signer.FetchGroupPublicKey("http://localhost:1323/group-public-key")
	if err != nil {
		t.Fatal(err)
	}

	_, err = signer.Sign([]byte("aaa"), 2, s1, gpk2, "http://localhost:1323/verify")
	if err != nil {
		t.Fatal(err)
	}

	passDay()
	rlBuf, err := signer.FetchRevocationList(s1, "http://localhost:1323/revocation-list")
	if err != nil {
		t.Fatal(err)
	}

	s12, err := signer.RequestUpdate(s1, rlBuf, gpk2, "http://localhost:1323/update-credential")
	if err != nil {
		t.Fatal(err)
	}

	_, err = signer.RequestUpdate(s2, rlBuf, gpk2, "http://localhost:1323/update-credential")
	if err != nil {
		t.Fatal(err)
	}

	ns, err = signer.RequestUpdate(s1, rlBuf, gpk2, "http://localhost:1323/update-credential")
	if err != nil {
		t.Fatal(ns, err)
	}

	if !bytes.Equal(s12, ns) {
		t.Fatal(ns, s12, ns)
	}

	passDay()
	rlBuf3, err := signer.FetchRevocationList(s12, "http://localhost:1323/revocation-list")
	if err != nil {
		t.Fatal(err)
	}

	_, err = signer.RequestUpdate(s12, rlBuf3, gpk2, "http://localhost:1323/update-credential")
	if err != nil {
		t.Fatal(err)
	}
}
