package quiz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientViewOfQuestion(t *testing.T) {
	item := QuizItem{
		Question: "What is the capital of Japan?",
		Options:  []string{"Tokyo", "Kyoto", "Osaka", "Nagoya"},
		Answer:   0,
	}

	output := item.ClientView()

	assert.Contains(t, output, "Question: What is the capital of Japan?")
	assert.Contains(t, output, "A) Tokyo")
	assert.Contains(t, output, "B) Kyoto")
	assert.Contains(t, output, "C) Osaka")
	assert.Contains(t, output, "D) Nagoya")
	assert.NotContains(t, output, item.Answer)
}
