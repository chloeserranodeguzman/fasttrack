package quiz

import "fmt"

type Scorer struct {
	CorrectAnswers int
	TotalQuestions int
}

func (s *Scorer) Evaluate(selectedOption int, correctOption int) {
	s.TotalQuestions++
	if selectedOption == correctOption {
		s.CorrectAnswers++
	}
}

func (s Scorer) GetScore() string {
	return fmt.Sprintf("Your score: %d/%d\n", s.CorrectAnswers, s.TotalQuestions)
}
