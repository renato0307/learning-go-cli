package cmd

import (
	"github.com/renato0307/learning-go-cli/internal/config"
	"github.com/spf13/cobra"
)

// NewConfigureCommand creates the the configure command
func NewConfigureCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "configure",
		Short: "Configures the CLI",
		Long:  `Allows to define the API endpoints and the client credentials`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return execute(cmd, args)
		},
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

// execute implements all the logic associated with this command.
func execute(cmd *cobra.Command, args []string) error {
	clientId, _ := cmd.Flags().GetString(config.ClientIdFlag)
	clientSecret, _ := cmd.Flags().GetString(config.ClientSecretFlag)
	apiEndpoint, _ := cmd.Flags().GetString(config.APIEndpointFlag)
	tokenEndpoint, _ := cmd.Flags().GetString(config.TokenEndpointFlag)

	return config.WriteAuthenticationConfig(
		clientId,
		clientSecret,
		apiEndpoint,
		tokenEndpoint,
	)
}
