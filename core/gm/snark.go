package gm

import (
	gnarkprecomputes "github.com/akakou/gnark-precomputes"
	"github.com/akakou/zk-ban-system/core/gm/ent"
	"github.com/akakou/zk-ban-system/core/gm/ent/revocation"
	"github.com/akakou/zk-ban/highlevel"
)

var UpdateRequestVerifyingKeys []gnarkprecomputes.PreparableCircuit

type Witness = any
type Prepared = any
type Proof = any

type PreparableSnarkVerifierKey struct {
	VerifyingKey []byte
	Size         *RevocationListSize
}

var PreparableSnarkVerifierKeys = []*PreparableSnarkVerifierKey{}

type PreparedSnarkVerifier struct {
	VerifierKey *PreparableSnarkVerifierKey
	Prepared    []byte
}

type DoubleMapKey struct {
	First, Second int
}

var PreparedSnarkVerifiers map[DoubleMapKey]*PreparedSnarkVerifier = make(map[DoubleMapKey]*PreparedSnarkVerifier)

func (gm *GroupManager[T]) BuildVerifier(before, after int) (*PreparedSnarkVerifier, error) {
	var v []struct {
		CountAll     int `json:"count_all"`
		Count        int `json:"count"`
		SignedPeriod int `json:"signed_period"`
	}

	query := gm.DB.Client.Revocation.Query().
		Where(revocation.And(
			revocation.RevokedPeriodGTE(before),
			revocation.SignedPeriodLTE(before),
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
	size := RevocationListSize{
		NymsNumberPerSession: maxNym,
		SessionNumber:        sessionNum,
	}

	rlDB := query.AllX(*gm.DB.Ctx)

	key := SelectVerifierKeyFromSize(&size)
	rlWit := TranslateRLFromDBToWit(rlDB, key.Size)

	prepared, err := highlevel.PrepareVerification(rlWit, key.VerifyingKey)
	if err != nil {
		return nil, err
	}

	verifier := &PreparedSnarkVerifier{
		VerifierKey: key,
		Prepared:    prepared,
	}

	return verifier, nil

}
