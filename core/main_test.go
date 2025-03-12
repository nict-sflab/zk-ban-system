package core

import (
	"encoding/json"
	"testing"

	"github.com/akakou/zk-ban/circuit"
	"github.com/akakou/zk-ban/witness"
)

func TestAll(t *testing.T) {
	joinProverBuf, joinVerifyKey := Prepare(&circuit.JoinRequestCircuit{}, t)
	signProverBuf, signVerifyKey := Prepare(&circuit.SignCircuit{}, t)
	updateProverBuf, updateVerifyKey := Prepare(&circuit.UpdateCircuit{}, t)

	rl := witness.RevocationList{}
	rlBuf, err := json.Marshal(rl)
	if err != nil {
		t.Fatal(err)
	}

	gsk, gpk, err := witness.RandomGroupKeyPair()
	if err != nil {
		t.Fatal(err)
	}

	req, signer, err := RequestJoin(joinProverBuf, "")
	if err != nil {
		t.Fatal(err)
	}

	cred, err := IssueCredential[string](req, gsk.Bytes(), joinVerifyKey)
	if err != nil {
		t.Fatal(err)
	}

	signer, err = SetCredential(cred, signer)
	if err != nil {
		t.Fatal(err)
	}

	m := []byte("test")
	signature, err := Sign(m, 0, signer, gsk.Bytes(), signProverBuf)
	if err != nil {
		t.Fatal(err)
	}

	err = Verify(signature, m, gpk.Bytes(), signVerifyKey)
	if err != nil {
		t.Fatal(err)
	}
	
	req, err = RequestUpdate(signer, rlBuf, gpk.Bytes(), updateProverBuf)
	if err != nil {
		t.Fatal(err)
	}

	cred, err = UpdateCredential(req, gsk.Bytes(), gpk.Bytes(), rl, updateVerifyKey)
	if err != nil {
		t.Fatal(err)
	}

	_, err = SetCredential(cred, signer)
	if err != nil {
		t.Fatal(err)
	}
}
