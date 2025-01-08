package quiz

import (
	"fmt"
	"sort"
)

var ScoreStore []int

type Scorer struct {
	CorrectAnswers int
	TotalQuestions int
}

func (s *Scorer) EvaluateAnswers(answers []int) {
	quizItems := getQuizItems()

	for i, answer := range answers {
		if i < len(quizItems) {
			s.evaluate(answer, quizItems[i].Answer)
		}
	}
}

func (s *Scorer) evaluate(selectedOption int, correctOption int) {
	s.TotalQuestions++
	if selectedOption == correctOption {
		s.CorrectAnswers++
	}
}

func (s *Scorer) calculatePercentile() int {
	if len(ScoreStore) == 0 {
		return 100
	}

	userScore := s.CorrectAnswers

	sortedScores := append([]int{}, ScoreStore...)
	sort.Ints(sortedScores)

	count := 0
	for _, score := range sortedScores {
		if score < userScore {
			count++
		}
	}

	percentile := int(float64(count) / float64(len(sortedScores)) * 100)

	ScoreStore = append(ScoreStore, userScore)
	sort.Ints(ScoreStore)

	return percentile
}

func (s Scorer) GetScore() string {
	percentile := s.calculatePercentile()
	return fmt.Sprintf("Your score: %d/%d\nYou were better than %d%% of all quizzers.\n",
		s.CorrectAnswers, s.TotalQuestions, percentile)
}
