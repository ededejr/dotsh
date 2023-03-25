package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(resetCmd)
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Remove all scripts and the API key",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		scriptDir, err := makeScriptDir()
		if err != nil {
			return err
		}

		// Remove all files in the script directory in parallel
		files, err := os.ReadDir(scriptDir)
		if err != nil {
			return err
		}
		var wg sync.WaitGroup
		wg.Add(len(files))
		errs := make(chan error, len(files))
		for _, f := range files {
			go func(filename string) {
				defer wg.Done()
				filePath := filepath.Join(scriptDir, filename)
				err := os.Remove(filePath)
				if err != nil {
					errs <- err
				}
			}(f.Name())
		}
		wg.Wait()
		close(errs)
		for err := range errs {
			if err != nil {
				return err
			}
		}

		// Remove script directory
		err = os.Remove(scriptDir)
		if err != nil {
			return err
		}

		fmt.Printf("Script directory '%s' has been reset\n", scriptDir)

		return nil
	},
}
