package tests

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/chloeserranodeguzman/fasttrack/cmd"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnCorrectWhenUserSelectsA(t *testing.T) {
	rootCmd := &cobra.Command{}
	cmd.AddQuizCommand(rootCmd)

	userInput := bytes.NewBufferString("A\n")

	output := bytes.NewBufferString("")
	rootCmd.SetOut(output)
	rootCmd.SetIn(userInput)
	rootCmd.SetArgs([]string{"quiz"})

	rootCmd.Execute()

	assert.Contains(t, output.String(), "Correct!")
}

func TestShouldReturnIncorrectAndGiveTheRightAnswerWhenUserSelectsB(t *testing.T) {
	rootCmd := &cobra.Command{}
	cmd.AddQuizCommand(rootCmd)

	userInput := bytes.NewBufferString("B\n")

	output := bytes.NewBufferString("")
	rootCmd.SetOut(output)
	rootCmd.SetIn(userInput)
	rootCmd.SetArgs([]string{"quiz"})

	rootCmd.Execute()

	assert.Contains(t, output.String(), "Incorrect. The correct answer was: Tokyo")
}

func TestShouldReturnInvalidInputAndPromptToTryAgain(t *testing.T) {
	rootCmd := &cobra.Command{}
	cmd.AddQuizCommand(rootCmd)

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
	cmd.AddQuizCommand(rootCmd)

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

	assert.Contains(t, output, "Quiz complete!")
}

func TestShouldScoreTwoWhenOnlyFirstTwoAnswersCorrect(t *testing.T) {
	rootCmd := &cobra.Command{}
	cmd.AddQuizCommand(rootCmd)

	var stdout, stdin bytes.Buffer

	stdin.WriteString("A\nB\nC\nA\n")

	rootCmd.SetIn(&stdin)
	rootCmd.SetOut(&stdout)
	rootCmd.SetArgs([]string{"quiz"})

	rootCmd.Execute()

	output := stdout.String()

	assert.Contains(t, output, "Your score: 2/4")
}
