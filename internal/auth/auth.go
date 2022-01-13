package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/renato0307/learning-go-cli/internal/config"
)

// AccessToken represents an OAuth2 access token obtained using the client
// credentials flow
type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// NewAccessToken fetches a new access token from the OAuth2 server
func NewAccessToken() (AccessToken, error) {

	accessToken := AccessToken{}

	clientId := config.GetString(config.ClientIdFlag)
	clientSecret := config.GetString(config.ClientSecretFlag)

	bodyContent := fmt.Sprintf(
		"grant_type=client_credentials&client_id=%s&scope=",
		clientId)
	body := strings.NewReader(bodyContent)

	tokenEndpoint := config.GetString(config.TokenEndpointFlag)
	request, err := http.NewRequest("POST", tokenEndpoint, body)
	if err != nil {
		return accessToken, err
	}

	clientIdAndSecret := fmt.Sprintf("%s:%s", clientId, clientSecret)
	credentials := base64.StdEncoding.EncodeToString([]byte(clientIdAndSecret))
	authHeader := fmt.Sprintf("Basic %s", credentials)
	request.Header = map[string][]string{
		"Authorization": {authHeader},
		"Content-Type":  {"application/x-www-form-urlencoded"},
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return accessToken, err
	}

	responseContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return accessToken, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return accessToken, fmt.Errorf("error getting token: %s", responseContent)
	}

	err = json.Unmarshal(responseContent, &accessToken)

	return accessToken, err
}
