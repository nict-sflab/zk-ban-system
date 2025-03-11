package core

import (
	"encoding/json"
	"testing"

	"github.com/akakou/zk-ban/circuit"
	"github.com/akakou/zk-ban/highlevel"
	"github.com/akakou/zk-ban/snark"
	"github.com/akakou/zk-ban/witness"
)

func TestAll(t *testing.T) {
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

	req, err := RequestJoin(proverBuf, "")
	if err != nil {
		t.Fatal(err)
	}

	_, err = IssueCredential[string](req, gsk.Bytes(), verifyKey)
	if err != nil {
		t.Fatal(err)
	}
}
