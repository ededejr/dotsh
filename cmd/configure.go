package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/howeyc/gopass"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(configureCmd)
}

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure API key for OpenAI",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Prompt user for API key
		apiKey, err := gopass.GetPasswdPrompt("Enter your OpenAI API key: ", false, os.Stdin, os.Stdout)
		if err != nil {
			return err
		}

		// Convert apiKey to string
		apiKeyStr := string(apiKey)

		// Trim whitespace from API key
		apiKeyStr = strings.TrimSpace(apiKeyStr)

		// Create directory for storing API key if it doesn't exist
		scriptDir, err := makeScriptDir()
		if err != nil {
			return err
		}

		// Write API key to file
		apiKeyFile := filepath.Join(scriptDir, apiKeyFileName)
		err = os.WriteFile(apiKeyFile, []byte(apiKeyStr), 0600)
		if err != nil {
			return err
		}

		fmt.Println("API key configured successfully")

		return nil
	},
}
