package gm

import (
	"encoding/json"

	"github.com/akakou/zk-ban-system/core/core"
	corecore "github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/load"
	zkbanw "github.com/akakou/zk-ban/witness"
)

type GroupManager[T any] struct {
	GroupSecretKey         zkbanw.GroupSecretKey
	GroupPublicKey         zkbanw.GroupPublicKey
	JoinVerifyKey          corecore.VerifyingKey
	VerifierKeys           corecore.VerifyingKeys
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

	gskw, err := zkbanw.GroupSecretKeyFromBytes(gsk)
	if err != nil {
		return nil, err
	}

	gpk := gskw.Public()

	var joinVerifierKey core.VerifyingKey
	err = json.Unmarshal(load.JoinVerifierKey, joinVerifierKey)
	if err != nil {
		return nil, err
	}

	g := GroupManager[T]{
		GroupSecretKey:         zkbanw.GroupSecretKey{gskw},
		GroupPublicKey:         zkbanw.GroupPublicKey{gpk},
		JoinVerifyKey:          joinVerifierKey,
		VerifierKeys:           u,
		PreparedSnarkVerifiers: make(map[corecore.KeyIndex]*PreparedSnarkVerifier),
		DB:                     db,
	}

	return &g, nil
}
