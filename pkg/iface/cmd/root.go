package cmd

import (
	"github.com/spf13/cobra"
)

var (
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "gopenapi",
		Short: "Utitlity tools for OpenAPI implementing by Go",
	}
)

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}
