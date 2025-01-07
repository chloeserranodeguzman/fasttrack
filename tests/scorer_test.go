package tests

import (
	"testing"

	"github.com/chloeserranodeguzman/fasttrack/quiz"
	"github.com/stretchr/testify/assert"
)

func TestScorer(t *testing.T) {
	scorer := &quiz.Scorer{}

	scorer.Evaluate(0, 0)
	scorer.Evaluate(1, 1)
	scorer.Evaluate(2, 0)

	assert.Equal(t, 2, scorer.CorrectAnswers)
	assert.Equal(t, 3, scorer.TotalQuestions)
	assert.Contains(t, scorer.GetScore(), "Your score: 2/3")
}
