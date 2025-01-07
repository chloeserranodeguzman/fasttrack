package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/chloeserranodeguzman/fasttrack/quiz"
	"github.com/spf13/cobra"
)

var quizCmd = &cobra.Command{
	Use:   "quiz",
	Short: "Start a new quiz",
	Run:   startQuiz,
}

var HttpGet = http.Get
var HttpPost = http.Post

func startQuiz(cmd *cobra.Command, args []string) {
	resp, err := HttpGet("http://localhost:8080/questions")
	if err != nil {
		fmt.Println("Failed to fetch questions from server:", err)
		return
	}
	defer resp.Body.Close()

	var quizItems []quiz.QuizItem
	json.NewDecoder(resp.Body).Decode(&quizItems)

	reader := bufio.NewReader(cmd.InOrStdin())
	var userAnswers []int

	for _, item := range quizItems {
		displayQuestion(cmd, item)

		answer, exit := getAnswer(cmd, reader)
		if exit {
			fmt.Fprint(cmd.OutOrStdout(), "Exiting quiz...\n")
			return
		}

		userAnswerIndex := map[string]int{"A": 0, "B": 1, "C": 2, "D": 3}[answer]
		userAnswers = append(userAnswers, userAnswerIndex)
	}

	sendAnswersToServer(userAnswers)
}

func sendAnswersToServer(answers []int) {
	payload, _ := json.Marshal(map[string]interface{}{"answers": answers})
	resp, err := HttpPost("http://localhost:8080/answers", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Failed to submit answers:", err)
		return
	}
	defer resp.Body.Close()

	getQuizResults(resp.Body)
}

func getQuizResults(body io.Reader) {
	var result map[string]interface{}
	json.NewDecoder(body).Decode(&result)

	fmt.Printf("\nQuiz complete!\n%s\n", result["message"])
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
