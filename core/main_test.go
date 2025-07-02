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

	rl := witness.RevocationList{}
	rlBuf, err := json.Marshal(rl)
	if err != nil {
		t.Fatal(err)
	}

	gsk, gpk, err := witness.RandomGroupKeyPair()
	if err != nil {
		t.Fatal(err)
	}

	gmDB, err := gm.NewDB(&gm.DBConfig{
		Type:   "sqlite3",
		Config: "file::memory:?cache=shared&_fk=1",
	})

	if err != nil {
		t.Fatal(err)
	}

	g := gm.GroupManager[string]{
		GroupSecretKey:  gsk.Bytes(),
		GroupPublicKey:  gpk.Bytes(),
		JoinVerifyKey:   joinVerifyKey,
		UpdateVerifyKey: updateVerifyKey,
		DB:              gmDB,
	}

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

	updateReq, err := signer.RequestUpdate(rlBuf, s, gpk.Bytes(), updateProverBuf)
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
