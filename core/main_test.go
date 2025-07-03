package core

import (
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/core/gm"
	"github.com/akakou/zk-ban-system/core/signer"
	"github.com/akakou/zk-ban-system/core/verifier"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/witness"
)

var period = 1

func today() int64 {
	return int64(period)
}

func passDay() {
	period += 1
}

func TestAll(t *testing.T) {
	utils.Today = today
	before := today()

	joinProverBuf, joinVerifyKey, err := core.JoinRequestCircuit()
	if err != nil {
		t.Fatal(err)
	}

	signProverBuf, signVerifyKey, err := core.SignCircuit()
	if err != nil {
		t.Fatal(err)
	}

	updateProverBuf, updateVerifyKey, err := core.UpdateCircuit(0, 0)
	if err != nil {
		t.Fatal(err)
	}

	gsk, gpk, err := witness.RandomGroupKeyPair()
	if err != nil {
		t.Fatal(err)
	}

	g, err := gm.Default[string](&gm.DBConfig{
		Type:   "sqlite3",
		Config: "file::memory:?cache=shared&_fk=1",
	})

	if err != nil {
		t.Fatal(err)
	}

	g.GroupPublicKey = gpk.Bytes()
	g.GroupSecretKey = gsk.Bytes()
	g.JoinVerifyKey = joinVerifyKey

	g.VerifierKeys = append(g.VerifierKeys, &core.SnarkKey{
		VerifyingKey: updateVerifyKey,
		Size: &core.RevocationListSize{
			NymsNumberPerSession: 0,
			SessionNumber:        0,
		},
	})

	v := verifier.Verifier{
		GroupPublicKey: gpk.Bytes(),
		SignVerifyKey:  signVerifyKey,
	}

	reqBody, s, err := signer.RequestJoin(joinProverBuf, "")
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
	signature, err := signer.Sign(m, 0, s, gsk.Bytes(), signProverBuf)
	if err != nil {
		t.Fatal(err)
	}

	err = v.Verify(signature)
	if err != nil {
		t.Fatal(err)
	}

	passDay()

	rl, err := g.RevocationList(before)
	if err != nil {
		t.Fatal(err)
	}

	updateReq, err := signer.RequestUpdate(rl, s, gpk.Bytes(), updateProverBuf)
	if err != nil {
		t.Fatal(err)
	}

	cred, err = g.UpdateCredential(updateReq)
	if err != nil {
		t.Fatal(err)
	}

	failCred, err = g.UpdateCredential(updateReq)
	if err == nil {
		t.Fatal(failCred, err)
	}

	rawCred, err = base64.URLEncoding.DecodeString(cred)
	if err != nil {
		t.Fatal(err)
	}

	_, err = signer.SetCredential(rawCred, s)
	if err != nil {
		t.Fatal(err)
	}
}
