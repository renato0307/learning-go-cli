package auth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/renato0307/learning-go-cli/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNewAccessToken(t *testing.T) {

	testCases := []struct {
		Token      AccessToken
		Raw        string
		StatusCode int
		Purpose    string
		ErrorNil   bool
	}{
		{
			Token: AccessToken{
				AccessToken: "token",
				ExpiresIn:   1000,
				TokenType:   "Bearer",
			},
			StatusCode: 200,
			Purpose:    "success case",
			ErrorNil:   true,
		},
		{
			Token:      AccessToken{},
			StatusCode: 500,
			Purpose:    "get token failure case",
			ErrorNil:   false,
		},
		{
			Raw:        "this_is_invalid_json",
			StatusCode: 200,
			Purpose:    "invalid token",
			ErrorNil:   false,
		},
	}

	for _, tc := range testCases {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// to test the code sends an authorization token
			if r.Header.Get("Authorization") == "" {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Unauthorized"))
				return
			}

			// writes the response for other cases
			// we use the Raw field to form malformed responses
			w.WriteHeader(tc.StatusCode)
			if tc.Raw != "" {
				w.Write([]byte(tc.Raw))
			} else {
				body, _ := json.Marshal(tc.Token)
				w.Write(body)
			}
		}))
		defer srv.Close()

		config.Set(config.TokenEndpointFlag, srv.URL)

		token, err := NewAccessToken()

		if tc.ErrorNil {
			assert.NoError(t, err, "error found for "+tc.Purpose)
		} else {
			assert.Error(t, err, "error not found for "+tc.Purpose)
		}

		assert.Equal(t, tc.Token, token, "invalid token for "+tc.Purpose)
	}

}
