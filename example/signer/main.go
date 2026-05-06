package main

import (
	"time"

	"github.com/akakou/zk-ban-system/utils"
)

const SIGNER_PATH = "./signer.gob"
const GPK_PATH = "./gpk.bin"

const DEFAULT_KEY_PATH = "../../dump"

var KeyPath = ""

func main() {
	utils.PeriodUnit = time.Duration(time.Minute / 2)
	rootCmd.Execute()
}
