package core

import (
	"encoding/json"
	"testing"

	"github.com/akakou/zk-ban/highlevel"
	"github.com/akakou/zk-ban/snark"
	"github.com/consensys/gnark/frontend"
)

func Prepare[T frontend.Circuit](c T, t *testing.T) ([]byte, []byte) {
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
