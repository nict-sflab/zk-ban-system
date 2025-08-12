package gm

import (
	"fmt"

	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/core/gm/ent"
	"github.com/akakou/zk-ban/primitives"
	"github.com/akakou/zk-ban/witness"
)

func MakeRLWit(dbEntries []*ent.Revocation, rlSize *core.RevocationListSize, witSize *core.RevocationListSize) witness.RevocationList {
	rl := witness.EmptyRevocationList(witSize.NymsNumberPerSession)
	fmt.Printf("entries: %v\n", dbEntries)

	index := 0
	for t, nyms := range rlSize.NymsNumberPerSession {
		entry := dbEntries[index]
		tag := witness.SessionTag(
			int64(entry.Count),
			int64(entry.SignedPeriod))

		rl[t].SessionTag = tag

		for n := range nyms {
			entry := dbEntries[index]

			nymBig := primitives.NewBigInt(0)
			nymBig.SetBytes(entry.Nym)
			rl[t].Nyms[n] = nymBig

			index++
		}
	}
	return rl
}
