# Quiz App (CLI + REST API)

A simple CLI quiz app powered by a Golang REST API.

## Overview
The quiz app allows users to answer questions, submit answers, and receive scores. Users can compare their results with others who have taken the quiz.

## Acceptance Criteria
User stories/Use cases: 
- User should be able to get questions with a number of answers
- User should be able to select just one answer per question.
- User should be able to answer all the questions and then post his/hers answers and get back how many correct answers they had, displayed to the user.
- User should see how well they compared to others that have taken the quiz, eg. "You were better than 60% of all quizzers"

## Tech Stack
- **Backend:** Golang (REST API)
- **CLI:** Golang with [spf13/cobra](https://github.com/spf13/cobra)
- **Testing:** TDD (Test-Driven Development)

---

## Features
- Answer questions interactively.
- Submit answers and receive a score.
- Compare results to percentile of other users.

---

## Requirements
- Golang 1.20 or higher
- Git

---

## Installation
```bash
git clone https://github.com/chloeserranodeguzman/fasttrack.git
cd fasttrack
go mod tidy
```
---

## How to Run
- Start the API server
```
go run main.go serve
```

- Start CLI quiz
```
go run main.go quiz
```

- Run All Tests
```
go test ./...
```

---

## Testing
- client_test: Focus only on testing user input and interaction.
- sevrer_test: Focus only on testing endpoints.
- quiz_test: Focus only on that the answer is hidden from the user and formatting.
- scorer_test: Focus only on calculating percentile ranking and number of correct answers.
- helper_test: Focus only on answer validity and index mapping.

- Check test coverage
```
go test ./... -cover
```
