package core_test

import (
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/core/gm"
	"github.com/akakou/zk-ban-system/core/signer"
	"github.com/akakou/zk-ban-system/core/verifier"
	"github.com/akakou/zk-ban-system/load"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/witness"
)

var period = int64(1)

func today() int64 {
	return int64(period)
}

func passDay() {
	period += 1
}

func TestAll(t *testing.T) {
	utils.Today = today
	before := today()

	gsk, gpk, err := witness.RandomGroupKeyPair()
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

	verifierKey, err := load.DocodeVerifyingKey(load.SignVerifierKey)
	if err != nil {
		t.Fatal(err)
	}

	v := verifier.Verifier{
		GroupPublicKey: gpk,
		VerifyingKey:   verifierKey,
	}

	reqBody, s, err := signer.RequestJoin("")
	if err != nil {
		t.Fatal(err)
	}

	var req core.JoinRequest[string]
	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		t.Fatal(err)
	}

	cred, err := g.IssueCredential("", &req)
	if err != nil {
		t.Fatal(err)
	}

	failCred, err := g.IssueCredential("", &req)
	if err == nil {
		t.Fatal(failCred, err)
	}

	rawCred, err := base64.URLEncoding.DecodeString(cred)
	if err != nil {
		t.Fatal(err)
	}

	s, err = signer.SetCredential(rawCred, s)
	if err != nil {
		t.Fatal(err)
	}

	m := []byte("test")
	signatureBuf, err := signer.Sign(m, 0, s, gsk.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	var signature core.Signature
	err = json.Unmarshal(signatureBuf, &signature)
	if err != nil {
		t.Fatal(err)
	}

	err = v.Verify(&signature, period)
	if err != nil {
		t.Fatal(err)
	}

	passDay()

	rl, err := g.RevocationList(before)
	if err != nil {
		t.Fatal(err)
	}

	updateReqBuf, newSignerBuf, err := signer.RequestUpdate(rl, s, gpk.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	var updateReq core.UpdateRequest
	err = json.Unmarshal(updateReqBuf, &updateReq)
	if err != nil {
		t.Fatal(err)
	}

	cred, err = g.UpdateCredential(&updateReq)
	if err != nil {
		t.Fatal(err)
	}

	failCred, err = g.UpdateCredential(&updateReq)
	if err == nil {
		t.Fatal(failCred, err)
	}

	rawCred2, err := base64.URLEncoding.DecodeString(cred)
	if err != nil {
		t.Fatal(err)
	}

	newSignerBuf, err = signer.SetCredential(rawCred2, newSignerBuf)
	if err != nil {
		t.Fatal(err)
	}

	passDay()

	updateReq2Buf, _, err := signer.RequestUpdate(rl, newSignerBuf, gpk.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	var updateReq2 core.UpdateRequest
	err = json.Unmarshal(updateReq2Buf, &updateReq2)
	if err != nil {
		t.Fatal(err)
	}

	_, err = g.UpdateCredential(&updateReq2)
	if err != nil {
		t.Fatal(err)
	}
}
