package quiz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuizWithoutAnswersShouldNotHaveAnswer(t *testing.T) {
	item := QuizItem{
		Question: "What is the capital of Japan?",
		Options:  []string{"Tokyo", "Kyoto", "Osaka", "Nagoya"},
		Answer:   0,
	}

	output := item.GetQuizItemWithoutAnswers()

	assert.Contains(t, output, "Question: What is the capital of Japan?")
	assert.Contains(t, output, "A) Tokyo")
	assert.Contains(t, output, "B) Kyoto")
	assert.Contains(t, output, "C) Osaka")
	assert.Contains(t, output, "D) Nagoya")
	assert.NotContains(t, output, item.Answer)
}

func TestGetQuizWithoutAnswersShouldNotHaveAnswer(t *testing.T) {
	quizWithoutAnswers := GetQuizWithoutAnswers()

	assert.Equal(t, 4, len(quizWithoutAnswers))

	assert.Equal(t, "What is the capital of Japan?", quizWithoutAnswers[0]["question"])
	assert.ElementsMatch(t, []string{"Tokyo", "Kyoto", "Osaka", "Nagoya"}, quizWithoutAnswers[0]["options"])

	assert.Equal(t, "What is 2 + 2?", quizWithoutAnswers[1]["question"])
	assert.ElementsMatch(t, []string{"3", "4", "5", "6"}, quizWithoutAnswers[1]["options"])

	assert.NotContains(t, quizWithoutAnswers[0], "answer")
	assert.NotContains(t, quizWithoutAnswers[1], "answer")
}
