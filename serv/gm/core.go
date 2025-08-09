package gm

import (
	"github.com/akakou/zk-ban-system/core"
	coregm "github.com/akakou/zk-ban-system/core/gm"
)

type AuthToken[T any] func(*core.JoinRequest[T]) (string, error)
type GMServer[T any] struct {
	GM        *coregm.GroupManager[T]
	AuthToken AuthToken[T]
}
