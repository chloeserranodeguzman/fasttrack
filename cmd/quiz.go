package cmd

import (
	"bufio"
	"fmt"
	"strings"

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

		reader := bufio.NewReader(cmd.InOrStdin())
		fmt.Fprint(cmd.OutOrStdout(), "Enter your answer (A, B, C, D): ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		fmt.Fprintf(cmd.OutOrStdout(), "You selected: %s\n", answer)
	},
}

func AddQuizCommand(root *cobra.Command) {
	root.AddCommand(quizCmd)
}
