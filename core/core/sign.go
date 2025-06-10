package core

import "github.com/akakou/zk-ban/highlevel"

type Signature struct {
	Signature *highlevel.Signature
	Message   []byte
}
