package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/rakyll/openai-go"
	"github.com/rakyll/openai-go/chat"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(makeCmd)
}

var makeCmd = &cobra.Command{
	Use:   "make [name] [prompt]",
	Short: "Generate an executable script based on OpenAI chat completion",
	Long:  "Generate an executable script using OpenAI chat completion. To use this effectively just describe what you want the prompt to do.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		scriptFileName := fmt.Sprintf("%s.sh", dasherize(args[0]))
		input := args[1]

		apiKey, err := getAPIKey()
		if err != nil {
			return fmt.Errorf("failed to locate API key: %v", err)
		}

		// Initialize OpenAI client with API key
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
		defer cancel()
		s := openai.NewSession(apiKey)

		// Define prompt for generating script
		prompt := fmt.Sprintf("Write a bash script which does the following: \"%s\". I only want the code in a valid .sh file format.", input)

		// Generate script using OpenAI completion API
		client := chat.NewClient(s, "gpt-3.5-turbo")
		resp, err := client.CreateCompletion(ctx, &chat.CreateCompletionParams{
			Messages: []*chat.Message{
				{Role: "user", Content: prompt},
			},
		})

		if err != nil {
			return fmt.Errorf("failed to complete: %v", err)
		}

		scriptContent := resp.Choices[0].Message.Content
		scriptDir, err := makeScriptDir()

		if err != nil {
			return fmt.Errorf("failed to create script dir: %v", err)
		}

		// Write script file
		scriptFilePath := filepath.Join(scriptDir, scriptFileName)
		scriptFile, err := os.Create(scriptFilePath)
		if err != nil {
			return err
		}
		defer scriptFile.Close()
		_, err = scriptFile.WriteString(scriptContent)
		if err != nil {
			return err
		}

		// Make script executable
		err = os.Chmod(scriptFilePath, 0700)
		if err != nil {
			return err
		}

		fmt.Printf("Generated script '%s'\n", scriptFileName)

		return nil
	},
}
