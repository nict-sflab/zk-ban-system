package main

import (
	"github.com/akakou/zk-ban/load"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "import zk-snark key process in zk-ban",
	Run: func(cmd *cobra.Command, args []string) {
		load.ReDumpUserKey("", "join")
		load.ReDumpUserKey("", "sign")
		load.ReDumpUserKey("sample", "update")
	},
}
