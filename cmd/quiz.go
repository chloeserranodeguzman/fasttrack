package cmd

import (
	"fmt"

	"github.com/chloeserranodeguzman/fasttrack/quiz"
	"github.com/spf13/cobra"
)

var quizCmd = &cobra.Command{
	Use:   "quiz",
	Short: "Start a new quiz",
	Run: func(cmd *cobra.Command, args []string) {
		item := quiz.QuizItem{
			Question: "What is the capital of Japan?",
			Options:  []string{"Tokyo", "Kyoto", "Osaka", "Nagoya"},
			Answer:   0,
		}

		fmt.Fprint(cmd.OutOrStdout(), item.ClientView())
	},
}

func AddQuizCommand(root *cobra.Command) {
	root.AddCommand(quizCmd)
}
