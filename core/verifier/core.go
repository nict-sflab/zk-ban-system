package verifier

import (
	gnarkserializable "github.com/akakou/gnark-serializable"
	zkbanw "github.com/akakou/zk-ban/witness"
)

type Verifier struct {
	GroupPublicKey *zkbanw.GroupPublicKey
	VerifyingKey   *gnarkserializable.VerifyingKey
	CountMax       int64
}
