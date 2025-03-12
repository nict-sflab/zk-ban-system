package core

import (
	"encoding/json"
	"testing"

	"github.com/akakou/zk-ban/circuit"
	"github.com/akakou/zk-ban/highlevel"
	"github.com/akakou/zk-ban/snark"
	"github.com/akakou/zk-ban/witness"
	"github.com/consensys/gnark/frontend"
)

func prepare[T frontend.Circuit](c T, t *testing.T) ([]byte, []byte) {
	cc, err := snark.InitSNARK(c)
	if err != nil {
		t.Fatal(err)
	}

	encodedProveKey, err := snark.EncodeProverKey(cc.ProveKey)
	if err != nil {
		t.Fatal(err)
	}

	encodedVerifyKey, err := snark.EncodeVerifierKey(cc.VerifyKey)
	if err != nil {
		t.Fatal(err)
	}

	encodedCircuit, err := snark.EncodeCircuit(cc.ConstraintSystem)
	if err != nil {
		t.Fatal(err)
	}

	prover := highlevel.HighLevelSnarkProver{
		ConstraintSystem: encodedCircuit,
		ProveKey:         encodedProveKey,
	}

	proverBuf, err := json.Marshal(&prover)
	if err != nil {
		t.Fatal(err)
	}

	return proverBuf, encodedVerifyKey

}

func TestAll(t *testing.T) {
	joinProverBuf, joinVerifyKey := prepare(&circuit.JoinRequestCircuit{}, t)
	updateProverBuf, updateVerifyKey := prepare(&circuit.UpdateCircuit{}, t)

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

	req, err = RequestUpdate(signer, rlBuf, gpk.Bytes(), updateProverBuf)
	if err != nil {
		t.Fatal(err)
	}

	cred, err = UpdateCredential(req, signer, gsk.Bytes(), rl, updateVerifyKey)
	if err != nil {
		t.Fatal(err)
	}

	_, err = SetCredential(cred, signer)
	if err != nil {
		t.Fatal(err)
	}
}
