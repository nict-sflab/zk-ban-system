package main

import "github.com/akakou/zk-ban-system/dump"

func main() {
	dump.DumpBasicKeys()
	dump.DumpUpdateKeys(100, 10)
	dump.DumpUpdateKeys(200, 20)
	dump.DumpUpdateKeys(300, 30)
}
