package main

import (
	"time"

	"github.com/akakou/zk-ban-system/utils"
)

const SIGNER_PATH = "./signer.json"
const GPK_PATH = "./gpk.bin"

func main() {
	utils.PeriodUnit = time.Duration(time.Minute / 2)
	rootCmd.Execute()
}
