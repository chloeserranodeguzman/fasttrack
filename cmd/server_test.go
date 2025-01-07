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

func setupTestServer() *httptest.Server {
	mux := SetupRouter()
	return httptest.NewServer(mux)
}

func TestSubmitAnswersWithValidPayload(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	payload := `{"answers": [0, 1, 1, 1]}`
	resp, _ := http.Post(server.URL+"/answers", "application/json", bytes.NewBufferString(payload))
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)

	assert.Contains(t, response["message"], "Your score: 4/4")
}

func TestSubmitAnswersWithInvalidPayload(t *testing.T) {
	server := setupTestServer()
	defer server.Close()

	payload := `{"invalidField": [1, 2, 3]}`
	resp, _ := http.Post(server.URL+"/answers", "application/json", bytes.NewBufferString(payload))
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Contains(t, readBody(resp.Body), "Invalid request payload")
}

func readBody(body io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(body)
	return buf.String()
}
