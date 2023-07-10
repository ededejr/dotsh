package cmd

import (
	"fmt"
	"os"

	"github.com/ededejr/dotsh/build"
	"github.com/spf13/cobra"
)

var rootCmdFlags struct {
	Version   bool
	BuildInfo bool
}

func init() {
	rootCmd.Flags().BoolVarP(&rootCmdFlags.Version, "version", "v", false, "Get application version")
	rootCmd.Flags().BoolVarP(&rootCmdFlags.BuildInfo, "build-info", "b", false, "Get application build info")
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

var rootCmd = &cobra.Command{
	Use:   "dotsh",
	Short: "A CLI tool for generating and running executable scripts based on OpenAI chat completion",
	Run: func(cmd *cobra.Command, args []string) {
		hasBothFlagsPassed := rootCmdFlags.Version && rootCmdFlags.BuildInfo
		hasNoFlagsPassed := !rootCmdFlags.Version && !rootCmdFlags.BuildInfo

		if hasBothFlagsPassed || hasNoFlagsPassed {
			cmd.Help()
		}

		if rootCmdFlags.Version {
			fmt.Println(build.Version)
		}

		if rootCmdFlags.BuildInfo {
			fmt.Println("Version:", build.Version)
			fmt.Println("Commit:", build.Commit)
			fmt.Println("Date:", build.Date)
			fmt.Println("User:", build.User)
			fmt.Println("Built by:", build.BuiltBy)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
