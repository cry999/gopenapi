package cmd

import (
	"fmt"
	"os"

	"github.com/cry999/gopenapi/pkg/openapi"
	"github.com/spf13/cobra"
)

func init() {
	wd, _ := os.Getwd()

	bundleCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "openapi.yml", "bundled output file (default is openapi.yml)")
	bundleCmd.PersistentFlags().StringVarP(&projectDir, "project-dir", "p", wd, "")

	rootCmd.AddCommand(bundleCmd)
}

var (
	outputFile string

	bundleCmd = &cobra.Command{
		Use:   "bundle",
		Short: "Bundle files into one file, `openapy.yml`",
		RunE:  bundleRun,
	}
)

func bundleRun(cmd *cobra.Command, args []string) error {
	spec, err := openapi.LoadProject(projectDir)
	if err != nil {
		return fmt.Errorf("failed to load project '%s': %v", projectDir, err)
	}
	return openapi.DumpInOneFile(outputFile, spec)
}
