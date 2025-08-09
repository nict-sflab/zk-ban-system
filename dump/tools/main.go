package main

import "github.com/akakou/zk-ban-system/dump"

func main() {
	dump.DumpBasicKeys()
	// dump.DumpUpdateKeys(260, 540)
	// dump.DumpUpdateKeys(260, 270)
	// dump.DumpUpdateKeys(130, 540)
	// dump.DumpUpdateKeys(130, 270)
	// dump.DumpUpdateKeys(75, 270)
	// dump.DumpUpdateKeys(130, 135)
	// dump.DumpUpdateKeys(75, 135)
	// dump.DumpUpdateKeys(32, 135)
	dump.DumpUpdateKeys(5, 5)
}
