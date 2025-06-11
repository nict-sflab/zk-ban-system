package zkbansystem

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/akakou/zk-ban-system/client/signer"
	"github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/core/gm"
	coreverifier "github.com/akakou/zk-ban-system/core/verifier"
	gmserv "github.com/akakou/zk-ban-system/serv/gm"
	"github.com/akakou/zk-ban-system/serv/verifier"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/highlevel"
	"github.com/akakou/zk-ban/witness"
	"github.com/labstack/echo/v4"
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
	utils.Today = today

	gmDB, err := gm.NewDB(&gm.DBConfig{
		Type:   "sqlite3",
		Config: "file::memory:?cache=shared&_fk=1",
	})
	if err != nil {
		t.Fatal(err)
	}

	joinProver, joinVerify, err := core.JoinRequestCircuit()
	if err != nil {
		t.Fatal(err)
	}

	signProver, signVerifyKey, err := core.SignCircuit()
	if err != nil {
		t.Fatal(err)
	}

	updateProver, updateVerify, err := core.UpdateCircuit([]int32{})
	if err != nil {
		t.Fatal(err)
	}

	gsk, gpk, err := witness.RandomGroupKeyPair()
	if err != nil {
		t.Fatal(err)
	}

	g := gm.GroupManager[string]{
		GroupSecretKey:  gsk.Bytes(),
		GroupPublicKey:  gpk.Bytes(),
		JoinVerifyKey:   joinVerify,
		UpdateVerifyKey: updateVerify,
		DB:              gmDB,
	}

	gmServ := gmserv.GMServer[string]{
		GM: &g,
		AuthToken: func(t *core.JoinRequest[string]) (string, error) {
			return "token", nil
		},
	}

	e := echo.New()
	gmServ.SetupEchoServer(e)

	v := coreverifier.Verifier{
		GroupPublicKey: gpk.Bytes(),
		SignVerifyKey:  signVerifyKey,
	}

	verifierServ := verifier.VerifierServer{
		Verifier: &v,
	}

	verifierServ.SetupEchoServer(e)

	go func() {
		if err := e.Start(":1323"); err != nil {
			t.Fatal(err)
		}
	}()

	time.Sleep(SLEEP_TIME * time.Second)

	s, err := signer.RequestJoin("", joinProver, "http://localhost:1323/issue-credential")
	if err != nil {
		t.Fatal(err)
	}

	ns, err := signer.RequestJoin("", joinProver, "http://localhost:1323/issue-credential")
	if err == nil {
		t.Fatal(ns, err)
	}

	ss := highlevel.HighLevelSigner{}
	json.Unmarshal(s, &ss)

	gpk2, err := signer.FetchGroupPublicKey("http://localhost:1323/group-public-key")
	if err != nil {
		t.Fatal(err)
	}

	_, err = signer.Sign([]byte("aaa"), 0, s, gpk2, signProver, "http://localhost:1323/verify")
	if err != nil {
		t.Fatal(err)
	}

	passDay()
	rlBuf, err := signer.FetchRevocationList("http://localhost:1323/revocation-list")
	if err != nil {
		t.Fatal(err)
	}

	_, err = signer.RequestUpdate(s, rlBuf, gpk2, updateProver, "http://localhost:1323/update-credential")
	if err != nil {
		t.Fatal(err)
	}

	ns, err = signer.RequestUpdate(s, rlBuf, gpk2, updateProver, "http://localhost:1323/update-credential")
	if err == nil {
		t.Fatal(ns, err)
	}
}
