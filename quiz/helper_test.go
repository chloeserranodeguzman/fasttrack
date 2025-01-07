package quiz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAnswerIndex(t *testing.T) {
	helper := NewQuizHelper()

	assert.Equal(t, 0, helper.GetAnswerIndex("A"))
	assert.Equal(t, 1, helper.GetAnswerIndex("B"))
	assert.Equal(t, 2, helper.GetAnswerIndex("C"))
	assert.Equal(t, 3, helper.GetAnswerIndex("D"))
	assert.Equal(t, -1, helper.GetAnswerIndex("Z"))
}

func TestIsValidAnswer(t *testing.T) {
	helper := NewQuizHelper()

	assert.True(t, helper.IsValidAnswer("A"))
	assert.True(t, helper.IsValidAnswer("B"))
	assert.False(t, helper.IsValidAnswer("Z"))
	assert.False(t, helper.IsValidAnswer(""))
}
