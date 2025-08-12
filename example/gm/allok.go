package main

import "github.com/akakou/zk-ban-system/core"

func allOKAuth(t *core.JoinRequest[string]) (string, error) {
	return t.Option, nil
}
