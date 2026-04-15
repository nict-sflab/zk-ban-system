package gm

import (
	"fmt"
	"time"

	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/core/gm"
	"github.com/akakou/zk-ban-system/utils"
)

var PeriodRange = int64(365 * 2)
var TimeGranularity = int64(10)

func (serv *GMServer[T]) RunKeyPrecomputeDaemon() {
	waitTime := utils.PeriodUnit / time.Duration(TimeGranularity)
	last := int64(0)

	for {
		fmt.Println("period: ", utils.Period())
		time.Sleep(waitTime)
		period := utils.Period()
		if last >= period {
			continue
		}

		fmt.Printf("precomputes key ...")

		last = period
		nextPeriod := period + 1

		serv.GM.PreparedSnarkVerifiers = make(map[core.KeyIndex]*gm.PreparedUpdateSnarkVerifier)
		err := serv.GM.ReadyUpdateVerifyKeys(PeriodRange, nextPeriod)
		if err != nil {
			fmt.Printf("Failed to run key precompute daemon: ", err)
		}

		fmt.Printf("done")

	}
}
