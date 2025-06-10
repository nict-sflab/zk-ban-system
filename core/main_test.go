package core

import (
	"encoding/json"
	"testing"

	"github.com/akakou/zk-ban/witness"
)

func TestAll(t *testing.T) {
	joinProverBuf, joinVerifyKey, err := JoinRequestCircuit()
	if err != nil {
		t.Fatal(err)
	}

	signProverBuf, signVerifyKey, err := SignCircuit()
	if err != nil {
		t.Fatal(err)
	}

	updateProverBuf, updateVerifyKey, err := UpdateCircuit([]int32{})
	if err != nil {
		t.Fatal(err)
	}

	rl := witness.RevocationList{}
	rlBuf, err := json.Marshal(rl)
	if err != nil {
		t.Fatal(err)
	}

	gsk, gpk, err := witness.RandomGroupKeyPair()
	if err != nil {
		t.Fatal(err)
	}

	reqBody, signer, err := RequestJoin(joinProverBuf, "")
	if err != nil {
		t.Fatal(err)
	}

	var req JoinRequest[string]
	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		t.Fatal(err)
	}

	cred, err := IssueCredential(&req, gsk.Bytes(), joinVerifyKey)
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

	err = Verify(signature, gpk.Bytes(), signVerifyKey)
	if err != nil {
		t.Fatal(err)
	}

	updateReqBody, err := RequestUpdate(signer, rlBuf, gpk.Bytes(), updateProverBuf)
	if err != nil {
		t.Fatal(err)
	}

	var updateReq UpdateRequest
	err = json.Unmarshal(updateReqBody, &updateReq)
	if err != nil {
		t.Fatal(err)
	}

	cred, err = UpdateCredential(&updateReq, gsk.Bytes(), gpk.Bytes(), rl, updateVerifyKey)
	if err != nil {
		t.Fatal(err)
	}

	_, err = SetCredential(cred, signer)
	if err != nil {
		t.Fatal(err)
	}
}
