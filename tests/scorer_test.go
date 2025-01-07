package tests

import (
	"testing"

	"github.com/chloeserranodeguzman/fasttrack/quiz"
	"github.com/stretchr/testify/assert"
)

func TestScore(t *testing.T) {
	scorer := &quiz.Scorer{}

	scorer.Evaluate(0, 0)
	scorer.Evaluate(1, 1)
	scorer.Evaluate(2, 0)

	assert.Equal(t, 2, scorer.CorrectAnswers)
	assert.Equal(t, 3, scorer.TotalQuestions)
	assert.Contains(t, scorer.GetScore(), "Your score: 2/3")
}

func TestShouldScoreFiftyPercentWithPreviousScores(t *testing.T) {
	scorer := &quiz.Scorer{}

	previousScores := []int{1, 1, 2, 2, 2, 3, 3, 3, 4, 4}
	scorer.SetPreviousScores(previousScores)

	scorer.Evaluate(0, 0)
	scorer.Evaluate(1, 1)
	scorer.Evaluate(2, 0) // incorrect
	scorer.Evaluate(1, 1)

	percentile := scorer.CalculatePercentile()
	assert.Equal(t, 50, percentile, "Expected percentile to be 50")
}

func TestShouldScoreZeroPercentWithPreviousScores(t *testing.T) {
	scorer := &quiz.Scorer{}

	previousScores := []int{1, 1, 2, 2, 2, 3, 3, 3, 4, 4}
	scorer.SetPreviousScores(previousScores)

	scorer.Evaluate(0, 2)
	scorer.Evaluate(1, 0)
	scorer.Evaluate(2, 0)
	scorer.Evaluate(1, 2)

	percentile := scorer.CalculatePercentile()
	assert.Equal(t, 0, percentile, "Expected percentile to be 0")
}

func TestShouldScoreOneHundredPercentWithPreviousScores(t *testing.T) {
	scorer := &quiz.Scorer{}

	previousScores := []int{1, 1, 2, 2, 2, 3, 3, 3, 3, 3}
	scorer.SetPreviousScores(previousScores)

	scorer.Evaluate(0, 0)
	scorer.Evaluate(1, 1)
	scorer.Evaluate(2, 2)
	scorer.Evaluate(1, 1)

	percentile := scorer.CalculatePercentile()
	assert.Equal(t, 100, percentile, "Expected percentile to be 100")
}

func TestShouldScoreOneHundredPercentWithNoPreviousScores(t *testing.T) {
	scorer := &quiz.Scorer{}

	previousScores := []int{}
	scorer.SetPreviousScores(previousScores)

	scorer.Evaluate(0, 0)
	scorer.Evaluate(1, 1)
	scorer.Evaluate(2, 2)
	scorer.Evaluate(1, 1)

	percentile := scorer.CalculatePercentile()
	assert.Equal(t, 100, percentile, "Expected percentile to be 100")
}
