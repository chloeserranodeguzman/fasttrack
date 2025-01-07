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
	Run:   startQuiz,
}

func startQuiz(cmd *cobra.Command, args []string) {
	quizItems := quiz.GetQuestions()
	reader := bufio.NewReader(cmd.InOrStdin())

	for _, item := range quizItems {
		displayQuestion(cmd, item)

		answer, exit := getAnswer(cmd, reader)
		if exit {
			fmt.Fprint(cmd.OutOrStdout(), "Exiting quiz...\n")
			return
		}
		fmt.Fprintf(cmd.OutOrStdout(), "You selected: %s\n\n", answer)
	}

	fmt.Fprint(cmd.OutOrStdout(), "Quiz complete!\n")
}

func displayQuestion(cmd *cobra.Command, item quiz.QuizItem) {
	fmt.Fprint(cmd.OutOrStdout(), item.ClientView())
}

func getAnswer(cmd *cobra.Command, reader *bufio.Reader) (string, bool) {
	validOptions := map[string]bool{"A": true, "B": true, "C": true, "D": true}

	for {
		fmt.Fprint(cmd.OutOrStdout(), "Enter your answer (A, B, C, D): ")
		answer, err := reader.ReadString('\n')

		if err != nil {
			return "", true
		}

		answer = strings.TrimSpace(strings.ToUpper(answer))

		if validOptions[answer] {
			return answer, false
		}

		fmt.Fprint(cmd.OutOrStdout(), "Invalid input. Please enter A, B, C, or D.\n")
	}
}

func AddQuizCommand(root *cobra.Command) {
	root.AddCommand(quizCmd)
}
