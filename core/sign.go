package core

import (
	zkban "github.com/akakou/zk-ban"
)

type Signature struct {
	Signature *zkban.Signature
	Message   []byte
}
