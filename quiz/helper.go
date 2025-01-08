package quiz

type QuizHelper struct {
	answerMap    map[string]int
	validOptions map[string]bool
}

func NewQuizHelper() *QuizHelper {
	return &QuizHelper{
		answerMap:    map[string]int{"A": 0, "B": 1, "C": 2, "D": 3},
		validOptions: map[string]bool{"A": true, "B": true, "C": true, "D": true},
	}
}

func (qh *QuizHelper) GetAnswerIndex(answer string) int {
	index, exists := qh.answerMap[answer]
	if !exists {
		return -1
	}
	return index
}

func (qh *QuizHelper) IsValidAnswer(answer string) bool {
	return qh.validOptions[answer]
}
