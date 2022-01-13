package programming

import (
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
		RunE: func(cmd *cobra.Command, args []string) error {
			return executeProgrammingUuid(cmd, args)
		},
	}
}

// execute implements all the logic associated with this command.
// In this case as it is an aggregation command will return an error
func executeProgrammingUuid(cmd *cobra.Command, args []string) error {
	apiEndpoint := config.GetString(config.APIEndpointFlag)
	realUrl := fmt.Sprintf("%s/programming/uuid", apiEndpoint)
	request, err := http.NewRequest("POST", realUrl, nil)
	if err != nil {
		return fmt.Errorf("error creating the request to call the API: %s", err)
	}

	token, err := auth.NewAccessToken()
	if err != nil {
		return fmt.Errorf("error calling the JWT to call the API: %s", err)
	}
	request.Header = map[string][]string{
		"Authentication": {token.AccessToken},
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("error calling the API: %s", err)
	}
	defer response.Body.Close()

	uuid, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading the UUID: %s", err)
	}

	_, err = fmt.Println(string(uuid))
	if err != nil {
		return fmt.Errorf("error writting to the output: %s", err)
	}

	return nil
}
