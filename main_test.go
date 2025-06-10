package zkbansystem

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/gm"
	"github.com/akakou/zk-ban-system/signer"
	"github.com/akakou/zk-ban/highlevel"
	"github.com/akakou/zk-ban/witness"
	"github.com/labstack/echo/v4"
)

const SLEEP_TIME = 3

func TestMain(t *testing.T) {
	rl := witness.RevocationList{}
	rlBuf, err := json.Marshal(rl)
	if err != nil {
		t.Fatal(err)
	}

	joinProver, joinVerify, err := core.JoinRequestCircuit()
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

	g := gm.GroupManager[any]{
		GroupSecretKey:  gsk.Bytes(),
		GroupPublicKey:  gpk.Bytes(),
		JoinVerifyKey:   joinVerify,
		UpdateVerifyKey: updateVerify,
		CheckToken: func(t *core.JoinRequest[any]) error {
			return nil
		},
	}

	e := echo.New()
	g.SetupEchoServer(e)

	go func() {
		if err := e.Start(":1323"); err != nil {
			t.Fatal(err)
		}
	}()

	time.Sleep(SLEEP_TIME * time.Second)

	s, err := signer.RequestJoin([]byte{}, joinProver, "http://localhost:1323/issue-credential")
	if err != nil {
		t.Fatal(err)
	}

	ss := highlevel.HighLevelSigner{}
	json.Unmarshal(s, &ss)

	_, err = signer.RequestUpdate(s, rlBuf, gpk.Bytes(), updateProver, "http://localhost:1323/update-credential")
	if err != nil {
		t.Fatal(err)
	}
}
