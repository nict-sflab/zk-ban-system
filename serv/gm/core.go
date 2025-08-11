package gm

import (
	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/core/gm"
)

type AuthToken[T any] func(*core.JoinRequest[T]) (string, error)
type GMServer[T any] struct {
	GM        *gm.GroupManager[T]
	AuthToken AuthToken[T]
}
