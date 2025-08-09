package gm

import (
	"github.com/akakou/zk-ban-system/core"
	corecore "github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/core/gm/ent"
	"github.com/akakou/zk-ban-system/core/gm/ent/revocation"
)

func (gm *GroupManager[T]) QueryRL(before int64) (*core.RevocationList, error) {
	var v []struct {
		CountAll     int `json:"count_all"`
		Count        int `json:"count"`
		SignedPeriod int `json:"signed_period"`
	}

	gm.DB.Client.Revocation.Query().
		Where(revocation.And(
			revocation.RevokedPeriodGTE(int(before)),
			revocation.SignedPeriodLTE(int(before)),
		)).
		GroupBy(revocation.FieldSignedPeriod, revocation.FieldCount).
		Aggregate(ent.Count()).
		ScanX(*gm.DB.Ctx, &v)

	maxNym := 0
	for _, vv := range v {
		if vv.Count > maxNym {
			maxNym = vv.Count
		}
	}

	sessionNum := len(v)
	size := corecore.RevocationListSize{
		NymsNumberPerSession: maxNym,
		SessionNumber:        sessionNum,
	}

	rlDB := gm.DB.Client.Revocation.Query().
		Where(revocation.And(
			revocation.RevokedPeriodGTE(int(before)),
			revocation.SignedPeriodLTE(int(before)),
		)).
		Order(ent.Asc(revocation.FieldSignedPeriod, revocation.FieldCount)).
		AllX(*gm.DB.Ctx)

	rlSize := QueryProperRLWitSize(gm.VerifierKeys, &size)
	rlObj := TranslateRLFromDBToWit(rlDB, rlSize)

	return &core.RevocationList{
		List: &rlObj,
		Size: rlSize,
	}, nil
}
