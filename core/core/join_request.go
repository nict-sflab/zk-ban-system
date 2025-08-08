package core

import (
	zkban "github.com/akakou/zk-ban"
)

type JoinRequest[T any] struct {
	Period      int64             `json:"period"`
	JoinRequest zkban.JoinRequest `json:"join_request"`
	Option      T                 `json:"option"`
}
