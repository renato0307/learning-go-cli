package testhelpers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/renato0307/learning-go-cli/internal/auth"
)

// NewAuthTestServer create an httptest.Server to test command requiring
// API authentication
func NewAuthTestServer() *httptest.Server {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// to test the code sends an authorization token
		if r.Header.Get("Authorization") == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Unauthorized"))
			return
		}

		w.WriteHeader(http.StatusOK)
		body, _ := json.Marshal(auth.AccessToken{})
		w.Write(body)
	}))

	return srv
}

// NewAPITestServer create an httptest.Server to test command requiring
// to call the API
func NewAPITestServer(body string) *httptest.Server {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
	}))

	return srv
}
