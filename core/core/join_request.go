package core

import (
	zkban "github.com/akakou/zk-ban"
	zkbanw "github.com/akakou/zk-ban/witness"
)

type JoinRequest[T any] struct {
	Period        int64
	UserPublicKey zkbanw.UserPublicKey
	JoinRequest   *zkban.JoinRequest
	Option        T
}
