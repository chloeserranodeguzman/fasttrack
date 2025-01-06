package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quiz-app",
	Short: "CLI quiz application",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	AddQuizCommand(rootCmd)
	AddServeCommand(rootCmd)
}
