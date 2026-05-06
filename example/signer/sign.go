package main

import (
	"fmt"
	"log"

	"github.com/akakou/zk-ban-system/client/signer"
	"github.com/akakou/zk-ban/dump"
	"github.com/spf13/cobra"
)

var message string
var count int64
var verifierURL string

var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign process in zk-ban",
	Run: func(cmd *cobra.Command, args []string) {
		dump.KeyPath = KeyPath + "/"

		config, err := readFile(SIGNER_PATH)
		if err != nil {
			log.Fatal(err)
		}

		gpk, err := readFile(GPK_PATH)
		if err != nil {
			log.Fatal(err)
		}

		res, err := signer.Sign([]byte(message), count, config, gpk, verifierURL)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", res)
	},
}
