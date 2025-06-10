package gm

import "github.com/akakou/zk-ban-system/core"

type AuthToken[T any] func(*core.JoinRequest[T]) (string, error)

type GroupManager[T any] struct {
	GroupSecretKey  []byte
	GroupPublicKey  []byte
	JoinVerifyKey   []byte
	UpdateVerifyKey []byte
	AuthToken       AuthToken[T]
	DB              *DB
}
