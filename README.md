# Quiz App (CLI + REST API)

A simple CLI quiz app powered by a Golang REST API.

## Overview
The quiz app allows users to answer questions, submit answers, and receive scores. Users can compare their results with others who have taken the quiz.

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