package gm

import (
	"fmt"

	"github.com/akakou/zk-ban-system/core"
	corecore "github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/core/gm/ent"
	"github.com/akakou/zk-ban-system/core/gm/ent/revocation"
	"github.com/akakou/zk-ban/snark"
)

func (gm *GroupManager[T]) QueryRL(before, after int64) (*core.RevocationList, *snark.SizedSnarkVerifier, error) {
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
		GroupBy(revocation.FieldSignedPeriod).
		Aggregate(ent.Count()).
		ScanX(*gm.DB.Ctx, &v)

	nymNums := []int{}
	for _, vv := range v {
		nymNums = append(nymNums, vv.Count)
	}
	rlSize := corecore.RevocationListSize{
		NymsNumberPerPeriod: nymNums,
	}

	rlDB := gm.DB.Client.Revocation.Query().
		Where(revocation.And(
			revocation.RevokedPeriodGTE(int(before)),
			revocation.RevokedPeriodLT(int(after)),
			revocation.SignedPeriodLTE(int(before)),
		)).
		Order(ent.Asc(revocation.FieldSignedPeriod)).
		AllX(*gm.DB.Ctx)

	fmt.Printf("rl condition: %v <= revoked < %v & sign <= %v\nrl: %v\n", before, after, before, rlDB)

	verifier, err := SelectProperVerifier(gm.VerifierKeys, &rlSize)
	if err != nil {
		return nil, nil, err
	}

	wit := MakeRLWit(rlDB, &rlSize, &core.RevocationListSize{verifier.RLSize})

	return &core.RevocationList{
		List:    &wit,
		KeyName: verifier.Name,
	}, verifier, nil
}
