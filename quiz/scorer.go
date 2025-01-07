package quiz

import (
	"fmt"
	"sort"
)

type Scorer struct {
	CorrectAnswers int
	TotalQuestions int
	PreviousScores []int
}

func (s *Scorer) Evaluate(selectedOption int, correctOption int) {
	s.TotalQuestions++
	if selectedOption == correctOption {
		s.CorrectAnswers++
	}
}

func (s *Scorer) SetPreviousScores(scores []int) {
	s.PreviousScores = scores
}

func (s *Scorer) CalculatePercentile() int {
	if len(s.PreviousScores) == 0 {
		return 100 // If no previous scores, user is the best by default
	}

	userScore := s.CorrectAnswers

	sort.Ints(s.PreviousScores)

	count := 0
	for _, score := range s.PreviousScores {
		if score < userScore {
			count++
		}
	}

	percentile := int(float64(count) / float64(len(s.PreviousScores)) * 100)
	s.PreviousScores = append(s.PreviousScores, userScore)
	sort.Ints(s.PreviousScores)

	return percentile
}

func (s Scorer) GetScore() string {
	percentile := s.CalculatePercentile()
	return fmt.Sprintf("Your score: %d/%d\nYou were better than %d%% of all quizzers.\n",
		s.CorrectAnswers, s.TotalQuestions, percentile)
}
