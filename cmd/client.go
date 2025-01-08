package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/chloeserranodeguzman/fasttrack/quiz"
	"github.com/spf13/cobra"
)

var quizCmd = &cobra.Command{
	Use:   "quiz",
	Short: "Start a new quiz",
	Run:   StartQuiz,
}

var httpGet = http.Get
var httpPost = http.Post

func StartQuiz(cmd *cobra.Command, args []string) {
	quizHelper := quiz.NewQuizHelper()

	quizItems := getQuestionsFromServer()
	userAnswers := getAnswersFromUser(cmd, quizItems, quizHelper)
	results := getResultsFromServer(userAnswers)
	printResults(results)
}

func getQuestionsFromServer() []quiz.QuizItem {
	resp, err := httpGet("http://localhost:8080/questions")
	if err != nil {
		fmt.Println("Failed to fetch questions from server:", err)
		return nil
	}
	defer resp.Body.Close()

	var quizItems []quiz.QuizItem
	json.NewDecoder(resp.Body).Decode(&quizItems)
	return quizItems
}

func getAnswersFromUser(cmd *cobra.Command, quizItems []quiz.QuizItem, quizHelper *quiz.QuizHelper) []int {
	reader := bufio.NewReader(cmd.InOrStdin())
	var userAnswers []int

	for _, item := range quizItems {
		displayQuestion(cmd, item)

		answer, exit := getAnswer(cmd, reader, quizHelper)
		if exit {
			fmt.Fprint(cmd.OutOrStdout(), "Exiting quiz...\n")
			return nil
		}

		userAnswerIndex := quizHelper.GetAnswerIndex(answer)
		userAnswers = append(userAnswers, userAnswerIndex)
	}
	return userAnswers
}

func getResultsFromServer(answers []int) map[string]interface{} {
	payload, _ := json.Marshal(map[string]interface{}{"answers": answers})
	resp, err := httpPost("http://localhost:8080/answers", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Failed to submit answers:", err)
		return nil
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return result
}

func printResults(result map[string]interface{}) {
	if result == nil {
		fmt.Println("No results to display.")
		return
	}
	fmt.Printf("\nQuiz complete!\n%s\n", result["message"])
}

func displayQuestion(cmd *cobra.Command, item quiz.QuizItem) {
	fmt.Fprint(cmd.OutOrStdout(), item.GetQuizItemWithoutAnswers())
}

func getAnswer(cmd *cobra.Command, reader *bufio.Reader, quizHelper *quiz.QuizHelper) (string, bool) {
	for {
		fmt.Fprint(cmd.OutOrStdout(), "Enter your answer (A, B, C, D): ")
		answer, err := reader.ReadString('\n')

		if err != nil {
			return "", true
		}

		answer = strings.TrimSpace(strings.ToUpper(answer))
		if quizHelper.IsValidAnswer(answer) {
			return answer, false
		}

		fmt.Fprint(cmd.OutOrStdout(), "Invalid input. Please enter A, B, C, or D.\n")
	}
}

func AddQuizCommand(root *cobra.Command) {
	root.AddCommand(quizCmd)
}
