package gm

import (
	"math/big"

	"github.com/akakou/zk-ban-system/core/gm/ent"
	"github.com/akakou/zk-ban/witness"
)

func TranslateRLFromDBToWit(dbEntries []*ent.Revocation, size *RevocationListSize) witness.RevocationList {
	rl := witness.EmptyConstantRevocationAddList(size.SessionNumber, size.NymsNumberPerSession)
	tag := big.NewInt(0)

	tagIndex := 0
	nymIndex := 0

	for _, entry := range dbEntries {
		t := witness.SessionTag(
			big.NewInt(int64(entry.Count)),
			big.NewInt(int64(entry.SignedPeriod)))
		if tag.Cmp(tag) == 0 {
			n := big.NewInt(0)
			n.FillBytes(entry.Nym)

			rl[tagIndex].Nyms[nymIndex] = n
			nymIndex += 1
		} else {
			rl[tagIndex].SessionTag = t
			nymIndex = 0
			tagIndex += 1
		}
	}

	return rl
}
