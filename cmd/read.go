package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(readCmd)
}

var readCmd = &cobra.Command{
	Use:   "read [name]",
	Short: "Display the contents of a script",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		scriptName := fmt.Sprintf("%s.sh", dasherize(args[0]))
		scriptDir, err := makeScriptDir()
		if err != nil {
			return err
		}
		scriptPath := filepath.Join(scriptDir, scriptName)
		scriptContent, err := os.ReadFile(scriptPath)
		if err != nil {
			return err
		}
		fmt.Println(string(scriptContent))
		return nil
	},
}
