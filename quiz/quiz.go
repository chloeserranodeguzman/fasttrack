package quiz

type QuizItem struct {
	Question string
	Options  []string
	Answer   int
}

func getQuizItems() []QuizItem {
	return []QuizItem{
		{
			Question: "What is the capital of Japan?",
			Options:  []string{"Tokyo", "Kyoto", "Osaka", "Nagoya"},
			Answer:   0, // Correct answer: Tokyo (A)
		},
		{
			Question: "What is 2 + 2?",
			Options:  []string{"3", "4", "5", "6"},
			Answer:   1, // Correct answer: 4 (B)
		},
		{
			Question: "Which planet is known as the Red Planet?",
			Options:  []string{"Earth", "Mars", "Venus", "Jupiter"},
			Answer:   1, // Correct answer: Mars (B)
		},
		{
			Question: "Who wrote 'Hamlet'?",
			Options:  []string{"Charles Dickens", "William Shakespeare", "Mark Twain", "Leo Tolstoy"},
			Answer:   1, // Correct answer: Shakespeare (B)
		},
	}
}

func (q QuizItem) GetQuizItemWithoutAnswers() string {
	optionLabels := []string{"A", "B", "C", "D"}
	view := "Question: " + q.Question + "\n"
	for i, option := range q.Options {
		view += optionLabels[i] + ") " + option + "\n"
	}
	return view
}

func GetQuizWithoutAnswers() []map[string]interface{} {
	quizItems := getQuizItems()
	publicItems := make([]map[string]interface{}, len(quizItems))

	for i, item := range quizItems {
		publicItems[i] = map[string]interface{}{
			"question": item.Question,
			"options":  item.Options,
		}
	}
	return publicItems
}
