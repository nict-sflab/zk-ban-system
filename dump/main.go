package main

import (
	"github.com/akakou/zk-ban/dump"
	"github.com/akakou/zk-ban/witness"
)

func main() {
	path := "../keys/"
	dump.DumpBasicKeys(path)

	rl := witness.MakeGaussianRLSizeFromTotal(60, 1080, 60/4)
	dump.DumpUpdateKeys("sample", rl, path)
}
