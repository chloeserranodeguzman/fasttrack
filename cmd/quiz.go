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
		validOptions := map[string]bool{"A": true, "B": true, "C": true, "D": true}

		for {
			fmt.Fprint(cmd.OutOrStdout(), "Enter your answer (A, B, C, D): ")
			answer, err := reader.ReadString('\n')

			if err != nil {
				fmt.Fprint(cmd.OutOrStdout(), "No input received. Exiting...\n")
				break
			}

			answer = strings.TrimSpace(strings.ToUpper(answer))

			if validOptions[answer] {
				fmt.Fprintf(cmd.OutOrStdout(), "You selected: %s\n", answer)
				break
			} else {
				fmt.Fprint(cmd.OutOrStdout(), "Invalid input. Please enter A, B, C, or D.\n")
			}
		}
	},
}

func AddQuizCommand(root *cobra.Command) {
	root.AddCommand(quizCmd)
}
