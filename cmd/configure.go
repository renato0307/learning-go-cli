package cmd

import (
	"fmt"

	"github.com/renato0307/learning-go-cli/internal/config"
	"github.com/renato0307/learning-go-cli/internal/iostreams"
	"github.com/spf13/cobra"
)

// NewConfigureCommand creates the the configure command
func NewConfigureCommand(iostreams *iostreams.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Configures the CLI",
		Long:  `Allows to define the API endpoints and the client credentials`,
		RunE:  executeConfigure(iostreams),
	}

	cmd.Flags().StringP(config.ClientIdFlag,
		"c",
		"",
		"the client id to call the API")
	cmd.MarkFlagRequired(config.ClientIdFlag)

	cmd.Flags().StringP(config.ClientSecretFlag,
		"s",
		"",
		"the client secret to call the API")
	cmd.MarkFlagRequired(config.ClientSecretFlag)

	cmd.Flags().StringP(config.APIEndpointFlag,
		"a",
		"",
		"the API endpoint")
	cmd.MarkFlagRequired(config.APIEndpointFlag)

	cmd.Flags().StringP(config.TokenEndpointFlag,
		"t",
		"",
		"the endpoint to get authentication tokens")
	cmd.MarkFlagRequired(config.TokenEndpointFlag)

	return cmd
}

// executeConfigure implements all the logic associated with this command.
func executeConfigure(iostreams *iostreams.IOStreams) func(cmd *cobra.Command, args []string) error {

	return func(cmd *cobra.Command, args []string) error {
		clientId, _ := cmd.Flags().GetString(config.ClientIdFlag)
		clientSecret, _ := cmd.Flags().GetString(config.ClientSecretFlag)
		apiEndpoint, _ := cmd.Flags().GetString(config.APIEndpointFlag)
		tokenEndpoint, _ := cmd.Flags().GetString(config.TokenEndpointFlag)

		err := config.WriteAuthenticationConfig(
			clientId,
			clientSecret,
			apiEndpoint,
			tokenEndpoint,
		)

		if err == nil {
			fmt.Fprintf(iostreams.Out, "configuration updated!")
		}
		return err
	}
}
