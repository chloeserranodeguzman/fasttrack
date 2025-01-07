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
}
