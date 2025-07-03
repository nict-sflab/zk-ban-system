package gm

import (
	corecore "github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/core/gm/ent"
	"github.com/akakou/zk-ban-system/core/gm/ent/revocation"
	"github.com/akakou/zk-ban/highlevel"
)

func (gm *GroupManager[T]) QueryRLAndKey(before, after int64) (*highlevel.HighLevelRevocationList, *corecore.SnarkKey, error) {
	var v []struct {
		CountAll     int `json:"count_all"`
		Count        int `json:"count"`
		SignedPeriod int `json:"signed_period"`
	}

	query := gm.DB.Client.Revocation.Query().
		Where(revocation.And(
			revocation.RevokedPeriodGTE(int(before)),
			revocation.SignedPeriodLTE(int(before)),
		))

	query.GroupBy(revocation.FieldSignedPeriod, revocation.FieldCount).
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

	rlDB := query.AllX(*gm.DB.Ctx)

	key := QueryProperVerifier(gm.VerifierKeys, &size)
	rlObj := TranslateRLFromDBToWit(rlDB, key.Size)

	rl := highlevel.HighLevelRevocationList{}
	rl.FromRevocationList(&rlObj)

	return &rl, key, nil
}
