package gm

import corecore "github.com/akakou/zk-ban-system/core/core"

type GroupManager[T any] struct {
	GroupSecretKey         []byte
	GroupPublicKey         []byte
	JoinVerifyKey          []byte
	VerifierKeys           corecore.SnarkKeys
	PreparedSnarkVerifiers map[corecore.KeyIndex]*PreparedSnarkVerifier
	//  = make(map[corecore.KeyIndex]*PreparedSnarkVerifier)
	DB *DB
}

func Default[T any](dbConfig *DBConfig) (*GroupManager[T], error) {
	db, err := NewDB(dbConfig)
	if err != nil {
		return nil, err
	}

	g := GroupManager[T]{
		GroupSecretKey:         []byte{},
		GroupPublicKey:         []byte{},
		JoinVerifyKey:          []byte{},
		VerifierKeys:           make(corecore.SnarkKeys),
		PreparedSnarkVerifiers: make(map[corecore.KeyIndex]*PreparedSnarkVerifier),
		DB:                     db,
	}

	return &g, nil
}
