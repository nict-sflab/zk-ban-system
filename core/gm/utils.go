package gm

import (
	"math/big"

	corecore "github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/core/gm/ent"
	"github.com/akakou/zk-ban/witness"
)

func TranslateRLFromDBToWit(dbEntries []*ent.Revocation, size *corecore.RevocationListSize) witness.RevocationList {
	rl := witness.EmptyConstantRevocationAddList(size.SessionNumber, size.NymsNumberPerSession)
	tag := big.NewInt(0)

	tagIndex := 0
	nymIndex := 0

	for _, entry := range dbEntries {
		t := witness.SessionTag(
			int64(entry.Count),
			int64(entry.SignedPeriod))

		if tag.Cmp(&t.Int) == 0 {
			n := big.NewInt(0)
			n.SetBytes(entry.Nym)

			rl[tagIndex].Nyms[nymIndex] = n
			nymIndex += 1
		} else {
			rl[tagIndex].SessionTag = &t.Int
			nymIndex = 0
			tagIndex += 1
			tag = &t.Int
		}
	}

	return rl
}
