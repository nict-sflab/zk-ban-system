package gm

import (
	eddsa_bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381/twistededwards/eddsa"

	corecore "github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/load"
)

type GroupManager[T any] struct {
	GroupSecretKey         []byte
	GroupPublicKey         []byte
	JoinVerifyKey          []byte
	VerifierKeys           corecore.SnarkKeys
	PreparedSnarkVerifiers map[corecore.KeyIndex]*PreparedSnarkVerifier
	DB                     *DB
}

func Default[T any](gsk []byte, dbConfig *DBConfig) (*GroupManager[T], error) {
	db, err := NewDB(dbConfig)
	if err != nil {
		return nil, err
	}

	u, err := load.LoadGroupManagerUpdateKey()
	if err != nil {
		return nil, err
	}

	gskw := eddsa_bls12381.PrivateKey{}
	_, err = gskw.SetBytes(gsk)
	if err != nil {
		return nil, err
	}

	gpk := gskw.Public()

	g := GroupManager[T]{
		GroupSecretKey:         gsk,
		GroupPublicKey:         gpk.Bytes(),
		JoinVerifyKey:          load.JoinVerifierKey,
		VerifierKeys:           u,
		PreparedSnarkVerifiers: make(map[corecore.KeyIndex]*PreparedSnarkVerifier),
		DB:                     db,
	}

	return &g, nil
}
