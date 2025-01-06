package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the API server to serve quiz questions",
	Long:  `Start a simple HTTP server that serves quiz questions and handles quiz submissions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting API server on port 8080...")

		http.HandleFunc("/questions", handleQuestions)

		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Printf("Failed to start server: %v\n", err)
		}
	},
}

func AddServeCommand(root *cobra.Command) {
	root.AddCommand(serveCmd)
}

// temporary
func handleQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`[
		{"id": 1, "text": "What is the capital of Japan?", "options": ["Tokyo", "Kyoto", "Osaka", "Nagoya"]},
		{"id": 2, "text": "What is 2 + 2?", "options": ["3", "4", "5", "6"]}
	]`))
}
