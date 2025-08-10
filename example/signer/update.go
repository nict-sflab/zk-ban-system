package main

import (
	"fmt"
	"log"

	"github.com/akakou/zk-ban-system/client/signer"
	"github.com/akakou/zk-ban-system/serv/gm"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update process in zk-ban",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := readFile(SIGNER_PATH)
		if err != nil {
			log.Fatal(err)
		}

		gpk, err := readFile(GPK_PATH)
		if err != nil {
			log.Fatal(err)
		}

		rl, err := signer.FetchRevocationList(config, gmBase+gm.REVOCATION_LIST_PATH)
		if err != nil {
			log.Fatal(err)
		}

		res, err := signer.RequestUpdate(config, rl, gpk, gmBase+gm.UPDATE_CREDENTIAL_PATH)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(string(res))
		writeFile(res, SIGNER_PATH)
	},
}
