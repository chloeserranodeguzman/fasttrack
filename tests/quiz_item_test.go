package tests

import (
	"testing"

	"github.com/chloeserranodeguzman/fasttrack/quiz"
	"github.com/stretchr/testify/assert"
)

func TestClientViewOfQuestion(t *testing.T) {
	item := quiz.QuizItem{
		Question: "What is the capital of Japan?",
		Options:  []string{"Tokyo", "Kyoto", "Osaka", "Nagoya"},
		Answer:   0, // Correct answer is index 0 (Tokyo)
	}

	output := item.ClientView()

	expected := `Question: What is the capital of Japan?
A) Tokyo
B) Kyoto
C) Osaka
D) Nagoya
`
	assert.Equal(t, expected, output, "Output should match the expected format.")
}
