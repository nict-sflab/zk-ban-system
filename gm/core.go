package gm

import "github.com/akakou/zk-ban-system/core"

type CheckToken[T any] func(*core.JoinRequest[T]) error

type GroupManager[T any] struct {
	GroupSecretKey  []byte
	GroupPublicKey  []byte
	JoinVerifyKey   []byte
	UpdateVerifyKey []byte
	CheckToken      CheckToken[T]
}
