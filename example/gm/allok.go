//go:build debug
// +build debug

package main

import (
	"fmt"

	"github.com/akakou/zk-ban-system/core"
)

func allOKAuth(t *core.JoinRequest[string]) (string, error) {
	return t.Option, nil
}

func authToken() func(t *core.JoinRequest[string]) (string, error) {
	fmt.Println("all ok mode")
	return allOKAuth
}
