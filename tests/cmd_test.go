package tests

import (
	"bytes"
	"testing"

	"github.com/chloeserranodeguzman/fasttrack/cmd"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestQuestionAndFeedbackShowsWhenYouRunQuiz(t *testing.T) {
	rootCmd := &cobra.Command{}
	cmd.AddQuizCommand(rootCmd)

	output := bytes.NewBufferString("")
	rootCmd.SetOut(output)
	rootCmd.SetArgs([]string{"quiz"})

	err := rootCmd.Execute()
	assert.NoError(t, err)

	assert.Contains(t, output.String(), "What is the capital of Japan?")
	assert.Contains(t, output.String(), "A) Tokyo")
	assert.Contains(t, output.String(), "B) Kyoto")
	assert.Contains(t, output.String(), "C) Osaka")
	assert.Contains(t, output.String(), "D) Nagoya")

	assert.Contains(t, output.String(), "You selected:")
}
