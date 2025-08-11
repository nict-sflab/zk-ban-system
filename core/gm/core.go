package gm

import (
	gnarkserializable "github.com/akakou/gnark-serializable"
	corecore "github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/core/gm/ent"
	"github.com/akakou/zk-ban-system/load"
	"github.com/akakou/zk-ban/witness"
	zkbanw "github.com/akakou/zk-ban/witness"
)

type GroupManager[T any] struct {
	GroupSecretKey         zkbanw.GroupSecretKey
	GroupPublicKey         zkbanw.GroupPublicKey
	JoinVerifyKey          *gnarkserializable.VerifyingKey
	VerifierKeys           corecore.VerifyingKeys
	PreparedSnarkVerifiers map[corecore.KeyIndex]*PreparedSnarkVerifier
	DB                     *DB
}

func Default[T any](gsk []byte, dbConfig *DBConfig) (*GroupManager[T], error) {
	db, err := NewDB(dbConfig)
	if err != nil {
		return nil, err
	}

	joinVerifierKey, err := load.DocodeVerifyingKey(load.JoinVerifierKey)
	if err != nil {
		return nil, err
	}

	updateVerifierKeys, err := load.LoadGroupManagerUpdateKey()
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

var JoinCallbackX = func(cred *witness.Credential, idToken *ent.IDToken, db *DB) {}
var UpdateCallbackX = func(cred *witness.Credential, updateTicket *ent.UpdateTicket, db *DB) {}
