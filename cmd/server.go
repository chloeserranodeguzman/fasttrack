package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chloeserranodeguzman/fasttrack/quiz"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the API server to serve quiz questions",
	Run: func(cmd *cobra.Command, args []string) {
		mux := SetupRouter()
		fmt.Println("Starting API server on port 8080...")
		err := http.ListenAndServe(":8080", mux)
		if err != nil {
			fmt.Printf("Failed to start server: %v\n", err)
		}
	},
}

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/questions", getQuestions)
	mux.HandleFunc("/answers", evaluateAnswers)
	return mux
}

func getQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	quiz := quiz.GetQuizItems()
	json.NewEncoder(w).Encode(quiz)
}

func evaluateAnswers(w http.ResponseWriter, r *http.Request) {
	var submission struct {
		Answers []int `json:"answers"`
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&submission)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	quizItems := quiz.GetQuizItems()
	scorer := &quiz.Scorer{}

	for i, answer := range submission.Answers {
		if i < len(quizItems) {
			scorer.Evaluate(answer, quizItems[i].Answer)
		}
	}

	result := map[string]interface{}{
		"message": scorer.GetScore(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func AddServeCommand(root *cobra.Command) {
	root.AddCommand(serveCmd)
}
