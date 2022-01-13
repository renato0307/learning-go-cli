package programming

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/renato0307/learning-go-cli/internal/auth"
	"github.com/renato0307/learning-go-cli/internal/config"
	"github.com/spf13/cobra"
)

// NewProgrammingCmd represents the programming command
func NewProgrammingUuidCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "uuid",
		Short: "Generates an UUID",
		Long:  `Generates an UUID, with or without hiphens.`,
		Run: func(cmd *cobra.Command, args []string) {
			executeProgrammingUuid(cmd, args)
		},
	}
}

// execute implements all the logic associated with this command.
// In this case as it is an aggregation command will return an error
func executeProgrammingUuid(cmd *cobra.Command, args []string) {
	apiEndpoint := config.GetString(config.APIEndpointFlag)
	realUrl := fmt.Sprintf("%s/programming/uuid", apiEndpoint)
	request, err := http.NewRequest("POST", realUrl, nil)
	if err != nil {
		cobra.CheckErr(fmt.Errorf("error creating the request to call the API: %s", err))
	}

	token, err := auth.NewAccessToken()
	if err != nil {
		cobra.CheckErr(fmt.Errorf("error calling the JWT to call the API: %s", err))
	}
	request.Header = map[string][]string{
		"Authentication": {token.AccessToken},
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		cobra.CheckErr(fmt.Errorf("error calling the API: %s", err))
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		apiError, err := ioutil.ReadAll(response.Body)
		if err != nil {
			cobra.CheckErr(fmt.Sprintf("error parsing API error: %s", err))
		}
		cobra.CheckErr(fmt.Sprintf("error calling the API: %s", apiError))
	}

	uuid, err := ioutil.ReadAll(response.Body)
	if err != nil {
		cobra.CheckErr(fmt.Errorf("error reading the UUID: %s", err))
	}

	var anyJson map[string]interface{}
	json.Unmarshal(uuid, &anyJson)
	output, _ := json.MarshalIndent(anyJson, "", "  ")

	_, err = fmt.Println(string(output))
	if err != nil {
		cobra.CheckErr(fmt.Errorf("error writting to the output: %s", err))
	}
}
