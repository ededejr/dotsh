package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all scripts in the script directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		scriptDir, err := makeScriptDir()
		if err != nil {
			return err
		}

		// List files in the script directory
		files, err := os.ReadDir(scriptDir)
		if err != nil {
			return err
		}

		// Print script names without the ".sh" extension
		fmt.Println("Scripts:")
		for _, file := range files {
			if !file.IsDir() && filepath.Ext(file.Name()) == ".sh" {
				scriptName := strings.TrimSuffix(file.Name(), ".sh")
				fmt.Println("-", scriptName)
			}
		}

		return nil
	},
}
