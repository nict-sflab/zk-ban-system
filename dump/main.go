package main

import (
	"github.com/akakou/zk-ban/dump"
	"github.com/akakou/zk-ban/witness"
)

func main() {
	dump.DumpBasicKeys(dump.DumpSafeKeys)

	rl := witness.MakeGaussianRLSizeFromTotal(30, 30000, 60/4)
	dump.DumpUpdateKeys("sample", rl, dump.DumpSafeKeys)
}
