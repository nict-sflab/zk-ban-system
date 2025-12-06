package verifier

import (
	zkbanw "github.com/akakou/zk-ban/witness"
	"github.com/consensys/gnark/backend/groth16"
)

type Verifier struct {
	GroupPublicKey *zkbanw.GroupPublicKey
	VerifyingKey   *groth16.VerifyingKey
	CountMax       int64
}
