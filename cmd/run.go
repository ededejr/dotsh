package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run [script-name] [args...]",
	Short: "Execute a script generated by 'dotsh make'",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		scriptName := dasherize(args[0])
		scriptDir, err := makeScriptDir()

		if err != nil {
			return fmt.Errorf("failed to create script dir: %v", err)
		}

		scriptFilePath := filepath.Join(scriptDir, fmt.Sprintf("%s.sh", scriptName))
		if _, err := os.Stat(scriptFilePath); os.IsNotExist(err) {
			return fmt.Errorf("script '%s' does not exist", scriptName)
		}

		sysCmdArgs := []string{"-c", fmt.Sprintf("%s %s", scriptFilePath, strings.Join(args[1:], "\" \""))}
		sysCmd := exec.Command("/bin/bash", sysCmdArgs...)
		sysCmd.Stdout = os.Stdout
		sysCmd.Stderr = os.Stderr
		sysCmd.Stdin = os.Stdin

		err = sysCmd.Run()
		if err != nil {
			return err
		}

		return nil
	},
}
