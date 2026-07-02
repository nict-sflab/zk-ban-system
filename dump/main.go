package main

import (
	"github.com/akakou/zk-ban/dump"
	"github.com/akakou/zk-ban/test"
)

func main() {
	dump.DumpBasicKeys(dump.DumpSafeKeys)

	rl := test.EmptyGaussianRevocationList(30, 30000)
	dump.DumpUpdateKeys("sample", rl.Sizes(), dump.DumpSafeKeys)
}
