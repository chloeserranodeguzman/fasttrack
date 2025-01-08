package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnInvalidInputAndPromptToTryAgain(t *testing.T) {
	server := mockServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"Question":"What is 2+2?","Options":["1","2","3","4"]}]`))
	})

	httpGet = func(url string) (*http.Response, error) {
		return http.Get(server.URL + "/questions")
	}

	cmd := &cobra.Command{}
	output := bytes.NewBufferString("")
	input := bytes.NewBufferString("X\n")

	cmd.SetOut(output)
	cmd.SetIn(input)

	StartQuiz(cmd, nil)

	assert.Contains(t, output.String(), "Invalid input. Please enter A, B, C, or D.")
	promptCount := strings.Count(output.String(), "Enter your answer (A, B, C, D):")
	assert.Equal(t, 2, promptCount, "Prompt should appear twice")
}

func TestShouldHaveAllQuestionsNonrepeatingWhenTwoValidInputsGiven(t *testing.T) {
	server := mockServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/questions" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`[{"Question":"What is 2+2?","Options":["1","2","3","4"]},
	                        {"Question":"What is the capital of Japan?","Options":["Tokyo","Osaka","Kyoto","Nagoya"]}]`))
		} else if r.URL.Path == "/answers" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"You scored 2/2"}`))
		}
	})

	httpGet = func(url string) (*http.Response, error) {
		return http.Get(server.URL + "/questions")
	}

	httpPost = func(url, contentType string, body io.Reader) (*http.Response, error) {
		return http.Post(server.URL+"/answers", contentType, body)
	}

	cmd := &cobra.Command{}
	var stdout, stdin bytes.Buffer
	stdin.WriteString("A\nB\n")

	cmd.SetIn(&stdin)
	cmd.SetOut(&stdout)

	StartQuiz(cmd, nil)

	output := stdout.String()
	expectedQuestions := []string{
		"What is the capital of Japan?",
		"What is 2+2?",
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
			Body:       io.NopCloser(bytes.NewBufferString(mockResponse)),
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

func mockServer(t *testing.T, handlerFunc http.HandlerFunc) *httptest.Server {
	server := httptest.NewServer(handlerFunc)
	t.Cleanup(func() { server.Close() })
	return server
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
