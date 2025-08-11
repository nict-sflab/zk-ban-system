package main

import (
	"fmt"
	"math/big"
	"time"

	"crypto/rand"

	"github.com/akakou/zk-ban-system/utils"
	"github.com/spf13/cobra"
)

var periodLength int64

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Update daemon in zk-ban",
	Run: func(cmd *cobra.Command, args []string) {
		utils.PeriodUnit = time.Duration(time.Minute / 2)
		for {
			r, err := rand.Int(rand.Reader, big.NewInt(int64(utils.PeriodUnit)))
			if err != nil {
				fmt.Println(err)
				continue
			}

			sleep := time.Duration(r.Int64())
			fmt.Printf("wait %s...\n", sleep)
			time.Sleep(sleep)
			err = update()
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	},
}
