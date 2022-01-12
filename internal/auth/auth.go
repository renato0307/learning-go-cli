package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

type accessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	IdToken     string `json:"id_token"`
}

func NewJWT() (string, error) {

	body := strings.NewReader("grant_type=client_credentials&scope=")

	tokenEndpoint := viper.Get(config.TokenEndpointFlag)
	request, err := http.NewRequest("POST", tokenEndpoint.(string), body)
	if err != nil {
		return "", err
	}

	clientId := viper.Get(config.ClientIdFlag)
	clientSecret := viper.Get(config.ClientSecretFlag)
	clientIdAndSecret := fmt.Sprintf("%s:%s", clientId, clientSecret)
	credentials := base64.StdEncoding.EncodeToString([]byte(clientIdAndSecret))
	authHeader := fmt.Sprintf("Bearer %s", credentials)

	request.Header = map[string][]string{
		"Authentication": {authHeader},
		"Content-Type":   {"application/x-www-form-urlencoded"},
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}

	responseContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	jwtResponse := accessToken{}
	err = json.Unmarshal(responseContent, &jwtResponse)
	if err != nil {
		return "", err
	}

	return jwtResponse.AccessToken, nil
}
