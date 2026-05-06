package main

import (
	"os"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "signer",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.signer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.AddCommand(importCmd)
	rootCmd.AddCommand(joinCmd)
	rootCmd.AddCommand(signCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(daemonCmd)

	importCmd.Flags().StringVar(&KeyPath, "keyPath", DEFAULT_KEY_PATH, "Path of proving keys in zk-SNAKR")
	joinCmd.Flags().StringVar(&KeyPath, "keyPath", DEFAULT_KEY_PATH, "Path of proving keys in zk-SNAKR")
	signCmd.Flags().StringVar(&KeyPath, "keyPath", DEFAULT_KEY_PATH, "Path of proving keys in zk-SNAKR")
	updateCmd.Flags().StringVar(&KeyPath, "keyPath", DEFAULT_KEY_PATH, "Path of proving keys in zk-SNAKR")
	daemonCmd.Flags().StringVar(&KeyPath, "keyPath", DEFAULT_KEY_PATH, "Path of proving keys in zk-SNAKR")

	joinCmd.Flags().StringVar(&idToken, "token", "", "ID Token of Firebase")
	joinCmd.Flags().StringVar(&gmBase, "base", "http://localhost:8080", "GM's base URL")
	joinCmd.MarkFlagRequired("token")

	signCmd.Flags().StringVar(&message, "message", "", "Message")
	signCmd.Flags().Int64Var(&count, "count", 0, "Count")
	signCmd.Flags().StringVar(&verifierURL, "url", "http://localhost:8000/verify", "Verifier's URL")
	signCmd.MarkFlagRequired("message")
	signCmd.MarkFlagRequired("count")

	updateCmd.Flags().StringVar(&gmBase, "base", "http://localhost:8080", "GM's base URL")

	daemonCmd.Flags().StringVar(&gmBase, "base", "http://localhost:8080", "GM's base URL")
	daemonCmd.Flags().Int64Var(&periodLength, "period_len", int64(time.Minute/2), "Period Length")
}
