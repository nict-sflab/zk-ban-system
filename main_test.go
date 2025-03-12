package zkbansystem

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/akakou/zk-ban-system/client"
	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/server"
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

	gm := server.GroupManager{
		GroupSecretKey:  gsk.Bytes(),
		GroupPublicKey:  gpk.Bytes(),
		JoinVerifyKey:   joinVerify,
		UpdateVerifyKey: updateVerify,
	}

	e := echo.New()
	issueCred := server.IssueCredential(&gm)
	updateCred := server.UpdateCredential(&gm)

	e.POST("/issue-credential", issueCred)
	e.POST("/update-credential", updateCred)

	go func() {
		if err := e.Start(":1323"); err != nil {
			t.Fatal(err)
		}
	}()

	time.Sleep(SLEEP_TIME * time.Second)

	signer, err := client.RequestJoin([]byte{}, joinProver, "http://localhost:1323/issue-credential")
	if err != nil {
		t.Fatal(err)
	}

	s := highlevel.HighLevelSigner{}
	json.Unmarshal(signer, &s)

	_, err = client.RequestUpdate(signer, rlBuf, gpk.Bytes(), updateProver, "http://localhost:1323/update-credential")
	if err != nil {
		t.Fatal(err)
	}
}
