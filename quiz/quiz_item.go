package quiz

import "fmt"

type QuizItem struct {
	Question string
	Options  []string
	Answer   int
}

func (q QuizItem) ClientView() string {
	optionLabels := []string{"A", "B", "C", "D"}
	output := fmt.Sprintf("Question: %s\n", q.Question)

	for i, option := range q.Options {
		output += fmt.Sprintf("%s) %s\n", optionLabels[i], option)
	}

	return output
}
