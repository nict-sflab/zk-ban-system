package main

import (
	"time"

	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/dump"
)

const SIGNER_PATH = "./signer.gob"
const GPK_PATH = "./gpk.bin"

func main() {
	dump.KeyPath = "../../dump/"
	utils.PeriodUnit = time.Duration(time.Minute / 2)
	rootCmd.Execute()
}
