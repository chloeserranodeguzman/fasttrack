package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubmitAnswersWithValidPayload(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	payload := `{"answers": [0, 1, 1, 1]}`
	resp, err := sendPostRequest(server.URL, "/answers", payload)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response map[string]interface{}
	readAndDecode(resp, &response)

	assert.Contains(t, response["message"], "Your score: 4/4")
}

func TestSubmitAnswersWithInvalidPayload(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	payload := `{"invalidField": [1, 2, 3]}`
	resp, err := sendPostRequest(server.URL, "/answers", payload)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Contains(t, readBodyAsString(resp), "Invalid request payload")
}

func TestGetQuestionsShouldHaveQuestionAndOptions(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	resp, err := sendGetRequest(server.URL, "/questions")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var questions []map[string]interface{}
	readAndDecode(resp, &questions)

	assert.NotEmpty(t, questions)
	assert.Contains(t, questions[0], "question")
	assert.Contains(t, questions[0], "options")
}

func setupTestServer() *httptest.Server {
	mux := setupRouter()
	return httptest.NewServer(mux)
}

func sendPostRequest(serverURL, endpoint, payload string) (*http.Response, error) {
	return http.Post(serverURL+endpoint, "application/json", bytes.NewBufferString(payload))
}

func sendGetRequest(serverURL, endpoint string) (*http.Response, error) {
	return http.Get(serverURL + endpoint)
}

func readAndDecode(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func readBodyAsString(resp *http.Response) string {
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body)
}
