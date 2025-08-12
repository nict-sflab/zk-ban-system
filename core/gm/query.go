package gm

import (
	"fmt"

	"github.com/akakou/zk-ban-system/core"
	corecore "github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/core/gm/ent"
	"github.com/akakou/zk-ban-system/core/gm/ent/revocation"
)

func (gm *GroupManager[T]) QueryRL(before, after int64) (*core.RevocationList, *core.SizedSnarkVerifier, error) {
	var v []struct {
		CountAll     int `json:"count_all"`
		Count        int `json:"count"`
		SignedPeriod int `json:"signed_period"`
	}

	gm.DB.Client.Revocation.Query().
		Where(revocation.And(
			revocation.RevokedPeriodGTE(int(before)),
			revocation.RevokedPeriodLT(int(after)),
			revocation.SignedPeriodLTE(int(before)),
		)).
		GroupBy(revocation.FieldSignedPeriod, revocation.FieldCount).
		Aggregate(ent.Count()).
		ScanX(*gm.DB.Ctx, &v)

	nymNums := []int{}
	for _, vv := range v {
		nymNums = append(nymNums, vv.Count)
	}
	rlSize := corecore.RevocationListSize{
		NymsNumberPerSession: nymNums,
	}

	rlDB := gm.DB.Client.Revocation.Query().
		Where(revocation.And(
			revocation.RevokedPeriodGTE(int(before)),
			revocation.RevokedPeriodLT(int(after)),
			revocation.SignedPeriodLTE(int(before)),
		)).
		Order(ent.Asc(revocation.FieldSignedPeriod, revocation.FieldCount)).
		AllX(*gm.DB.Ctx)

	fmt.Printf("rl condition: %v <= revoked < %v & sign <= %v\nrl: %v\n", before, after, before, rlDB)

	verifier, keyIndex, err := SelectProperVerifier(gm.VerifierKeys, &rlSize)
	if err != nil {
		return nil, nil, err
	}

	wit := MakeRLWit(rlDB, &rlSize, verifier.RLSize)

	return &core.RevocationList{
		List:  &wit,
		Index: keyIndex,
	}, verifier, nil
}
