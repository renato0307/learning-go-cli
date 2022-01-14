package programming

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/renato0307/learning-go-cli/internal/auth"
	"github.com/renato0307/learning-go-cli/internal/config"
	"github.com/renato0307/learning-go-cli/internal/iostreams"
	"github.com/spf13/cobra"
)

const NoHyphensFlag string = "no-hyphens"

// NewProgrammingCmd represents the programming command
func NewProgrammingUuidCmd(iostreams *iostreams.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uuid",
		Short: "Generates an UUID",
		Long:  `Generates an UUID, with or without hyphens.`,
		RunE:  executeProgrammingUuid(iostreams),
	}

	cmd.Flags().Bool(NoHyphensFlag,
		false,
		"if set the UUID generated will not contains hyphens")

	return cmd
}

// executeProgrammingUuid implements all the logic associated with this command.
// In this case as it is an aggregation command will return an error
func executeProgrammingUuid(iostreams *iostreams.IOStreams) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {

		// creates the base request
		apiEndpoint := config.GetString(config.APIEndpointFlag)
		realUrl := fmt.Sprintf("%s/programming/uuid", apiEndpoint)
		request, err := http.NewRequest("POST", realUrl, nil)
		if err != nil {
			return fmt.Errorf("error creating the request to call the API: %w", err)
		}

		// handles the "no-hyphens" flag
		noHyphens, err := cmd.Flags().GetBool(NoHyphensFlag)
		if err != nil {
			return err
		}
		if noHyphens {
			q := request.URL.Query()
			q.Add("no-hyphens", "true")
			request.URL.RawQuery = q.Encode()
		}

		// adds authentication
		token, err := auth.NewAccessToken()
		if err != nil {
			return fmt.Errorf("error getting the JWT to call the API: %w", err)
		}
		request.Header = map[string][]string{
			"Authentication": {token.AccessToken},
		}

		// calls API and reads response
		response, err := http.DefaultClient.Do(request)
		if err != nil {
			return fmt.Errorf("error calling the API: %w", err)
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			apiError, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return fmt.Errorf("error parsing API error: %w", err)
			}

			err = errors.New(string(apiError))
			return fmt.Errorf("error calling the API: %w", err)
		}

		uuid, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return fmt.Errorf("error reading the UUID: %w", err)
		}

		// parse and print response as indented JSON
		var anyJson map[string]interface{}
		err = json.Unmarshal(uuid, &anyJson)
		if err != nil {
			return fmt.Errorf("parsing API response: %w", err)
		}
		output, _ := json.MarshalIndent(anyJson, "", "  ")

		_, err = fmt.Fprintln(iostreams.Out, string(output))
		if err != nil {
			return fmt.Errorf("error writing to the output: %w", err)
		}

		return nil
	}
}
