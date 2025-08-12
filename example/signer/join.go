package main

import (
	"fmt"
	"log"

	"github.com/akakou/zk-ban-system/client/signer"
	"github.com/akakou/zk-ban-system/serv/gm"
	"github.com/spf13/cobra"
)

var idToken string
var gmBase string

var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Join process in zk-ban",
	Run: func(cmd *cobra.Command, args []string) {
		gpk, err := signer.FetchGroupPublicKey(gmBase + gm.GROUP_PUBLIC_KEY_PATH)
		if err != nil {
			log.Fatal(err)
		}

		config, err := signer.RequestJoin(idToken, gmBase+gm.ISSUE_CREDENTIAL_PATH)
		if err != nil {
			log.Fatal(err)
		}

		err = writeFile(gpk, GPK_PATH)
		if err != nil {
			log.Fatal(err)
		}
		writeFile(config, SIGNER_PATH)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", config)
	},
}
