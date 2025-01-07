package cmd

import (
	"bytes"
	"fmt"
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
