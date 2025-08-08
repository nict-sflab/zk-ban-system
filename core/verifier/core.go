package verifier

import (
	"github.com/akakou/zk-ban-system/core/core"
	zkbanw "github.com/akakou/zk-ban/witness"
)

type Verifier struct {
	GroupPublicKey *zkbanw.GroupPublicKey
	VerifyingKey   core.VerifyingKey
}
