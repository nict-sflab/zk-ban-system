package zkbansystem

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/akakou/zk-ban-system/client"
	"github.com/akakou/zk-ban-system/server"
	"github.com/akakou/zk-ban/circuit"
	"github.com/akakou/zk-ban/highlevel"
	"github.com/akakou/zk-ban/snark"
	"github.com/akakou/zk-ban/witness"
	"github.com/labstack/echo/v4"
)

const SLEEP_TIME = 5

func TestMain(t *testing.T) {
	joinSnark, err := snark.InitSNARK(&circuit.JoinRequestCircuit{})
	if err != nil {
		t.Fatal(err)
	}

	gsk, _, err := witness.RandomGroupKeyPair()
	if err != nil {
		t.Fatal(err)
	}

	proveKey, err := snark.EncodeProverKey(joinSnark.ProveKey)
	if err != nil {
		t.Fatal(err)
	}

	verifyKey, err := snark.EncodeVerifierKey(joinSnark.VerifyKey)
	if err != nil {
		t.Fatal(err)
	}

	circuit, err := snark.EncodeCircuit(joinSnark.ConstraintSystem)
	if err != nil {
		t.Fatal(err)
	}

	prover := highlevel.HighLevelSnarkProver{
		ConstraintSystem: circuit,
		ProveKey:         proveKey,
	}

	proverBuf, err := json.Marshal(&prover)
	if err != nil {
		t.Fatal(err)
	}

	gm := server.GroupManager{
		GroupSecretKey: gsk.Bytes(),
		VerifyKey:      verifyKey,
	}

	e := echo.New()
	issueCred := server.IssueCredential(&gm)

	e.POST("/issue-credential", issueCred)
	go func() {
		if err := e.Start(":1323"); err != nil {
			t.Fatal(err)
		}
	}()

	time.Sleep(SLEEP_TIME * time.Second)

	_, err = client.RequestJoin([]byte{}, proverBuf, "http://localhost:1323/issue-credential")
	if err != nil {
		t.Fatal(err)
	}
}
