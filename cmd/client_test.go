package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnInvalidInputAndPromptToTryAgain(t *testing.T) {
	rootCmd := &cobra.Command{}
	AddQuizCommand(rootCmd)

	userInput := bytes.NewBufferString("X\n")
	output := bytes.NewBufferString("")

	rootCmd.SetOut(output)
	rootCmd.SetIn(userInput)
	rootCmd.SetArgs([]string{"quiz"})

	rootCmd.Execute()

	assert.Contains(t, output.String(), "Invalid input. Please enter A, B, C, or D.")

	promptCount := strings.Count(output.String(), "Enter your answer (A, B, C, D):")

	assert.Equal(t, 2, promptCount, "Prompt should appear twice")
}

func TestShouldHaveAllQuestionsNonrepeatingWhenFourValidInputsGiven(t *testing.T) {
	rootCmd := &cobra.Command{}
	AddQuizCommand(rootCmd)

	var stdout, stdin bytes.Buffer

	stdin.WriteString("A\nB\nB\nB\n")

	rootCmd.SetIn(&stdin)
	rootCmd.SetOut(&stdout)
	rootCmd.SetArgs([]string{"quiz"})

	rootCmd.Execute()

	output := stdout.String()

	expectedQuestions := []string{
		"What is the capital of Japan?",
		"What is 2 + 2?",
		"Which planet is known as the Red Planet?",
		"Who wrote 'Hamlet'?",
	}

	for _, question := range expectedQuestions {
		count := strings.Count(output, question)
		assert.Equal(t, 1, count, fmt.Sprintf("Question '%s' appeared %d times", question, count))
	}
}

func TestGetQuestionsFromServerShouldSucceed(t *testing.T) {
	mockResponse := `[{"Question":"What is 2+2?","Options":["1","2","3","4"],"AnswerIndex":3}]`

	httpGet = func(url string) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(mockResponse)), // Use io.NopCloser
		}, nil
	}

	quizItems := getQuestionsFromServer()
	assert.NotNil(t, quizItems)
	assert.Equal(t, 1, len(quizItems))
	assert.Equal(t, "What is 2+2?", quizItems[0].Question)
}

func TestGetQuestionsFromServerShouldFail(t *testing.T) {
	httpGet = func(_ string) (*http.Response, error) {
		return nil, errors.New("server not reachable")
	}

	quizItems := getQuestionsFromServer()
	assert.Nil(t, quizItems)
}

func TestGetResultsFromServerShouldSucceed(t *testing.T) {
	mockResponse := `{"message":"You scored 4/4"}`

	httpPost = func(url, contentType string, body io.Reader) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(mockResponse)),
		}, nil
	}

	results := getResultsFromServer([]int{0, 1, 2})
	assert.NotNil(t, results)
	assert.Equal(t, "You scored 4/4", results["message"])
}

func TestGetResultsFromServerShouldFail(t *testing.T) {
	httpPost = func(_, _ string, _ io.Reader) (*http.Response, error) {
		return nil, errors.New("submission failed")
	}

	results := getResultsFromServer([]int{0, 1, 2})
	assert.Nil(t, results)
}

func TestPrintResultsShouldFail(t *testing.T) {
	output := captureStdout(func() {
		printResults(nil)
	})
	assert.Contains(t, output, "No results to display.")
}

func TestPrintResultsShouldSucceed(t *testing.T) {
	result := map[string]interface{}{"message": "Quiz complete!"}
	output := captureStdout(func() {
		printResults(result)
	})
	assert.Contains(t, output, "Quiz complete!")
}

func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = old

	return string(out)
}
