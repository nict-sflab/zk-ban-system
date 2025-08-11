package main

import (
	"fmt"
	"log"

	"github.com/akakou/zk-ban-system/client/signer"
	"github.com/akakou/zk-ban-system/serv/gm"
	"github.com/spf13/cobra"
)

func update() error {
	config, err := readFile(SIGNER_PATH)
	if err != nil {
		return err
	}

	gpk, err := readFile(GPK_PATH)
	if err != nil {
		return err
	}

	rl, err := signer.FetchRevocationList(config, gmBase+gm.REVOCATION_LIST_PATH)
	if err != nil {
		return err
	}

	res, err := signer.RequestUpdate(config, rl, gpk, gmBase+gm.UPDATE_CREDENTIAL_PATH)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", res)
	err = writeFile(res, SIGNER_PATH)
	if err != nil {
		return err
	}

	return nil
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update process in zk-ban",
	Run: func(cmd *cobra.Command, args []string) {
		err := update()
		if err != nil {
			log.Fatal(err)
		}
	},
}
