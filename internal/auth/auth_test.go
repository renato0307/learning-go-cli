package auth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/renato0307/learning-go-cli/internal/config"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestNewJWT(t *testing.T) {

	testCases := []struct {
		Token      accessToken
		StatusCode int
		Purpose    string
		ErrorNil   bool
	}{
		{
			Token: accessToken{
				AccessToken: "token",
				ExpiresIn:   1000,
				TokenType:   "Bearer",
				IdToken:     "id_token",
			},
			StatusCode: 200,
			Purpose:    "success case",
			ErrorNil:   true,
		},
		{
			Token:      accessToken{},
			StatusCode: 500,
			Purpose:    "get token failure case",
			ErrorNil:   false,
		},
		{
			Token:      accessToken{},
			StatusCode: 200,
			Purpose:    "invalid token",
			ErrorNil:   false,
		},
	}

	for _, tc := range testCases {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(tc.StatusCode)
			body, _ := json.Marshal(tc.Token)
			w.Write(body)
		}))
		defer srv.Close()

		viper.Set(config.TokenEndpointFlag, srv.URL)

		jwt, err := NewJWT()

		assert.Nil(t, err, tc.ErrorNil)
		assert.Equal(t, tc.Token.AccessToken, jwt)
	}

}
