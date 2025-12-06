package gm

import (
	corecore "github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban/load"
	zkbanw "github.com/akakou/zk-ban/witness"
	"github.com/consensys/gnark/backend/groth16"
)

type GroupManager[T any] struct {
	GroupSecretKey         zkbanw.GroupSecretKey
	GroupPublicKey         zkbanw.GroupPublicKey
	JoinVerifyKey          *groth16.VerifyingKey
	VerifierKeys           corecore.SizedVerifyingKeys
	PreparedSnarkVerifiers map[corecore.KeyIndex]*PreparedSnarkVerifier
	DB                     *DB
}

func Default[T any](gsk []byte, dbConfig *DBConfig) (*GroupManager[T], error) {
	db, err := NewDB(dbConfig)
	if err != nil {
		return nil, err
	}

	joinVerifierKey, err := load.LoadBasicGroupManagerKey("join")
	if err != nil {
		return nil, err
	}

	updateVerifierKeys, err := load.LoadGroupManagerUpdateKeys()
	if err != nil {
		return nil, err
	}

	gskw, err := zkbanw.GroupSecretKeyFromBytes(gsk)
	if err != nil {
		return nil, err
	}

	gpk := gskw.Public()

	g := GroupManager[T]{
		GroupSecretKey:         zkbanw.GroupSecretKey{gskw},
		GroupPublicKey:         zkbanw.GroupPublicKey{gpk},
		JoinVerifyKey:          joinVerifierKey,
		VerifierKeys:           updateVerifierKeys,
		PreparedSnarkVerifiers: make(map[corecore.KeyIndex]*PreparedSnarkVerifier),
		DB:                     db,
	}

	return &g, nil
}
