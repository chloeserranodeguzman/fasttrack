package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd is the base command
var rootCmd = &cobra.Command{
	Use:   "quiz-app",
	Short: "CLI quiz application",
}

// Execute adds all child commands to the root
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// Initialize and attach commands
func init() {
	AddQuizCommand(rootCmd)
	AddServeCommand(rootCmd) // Attach the serve command properly
}
