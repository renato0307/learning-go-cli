package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	ClientIdFlag      string = "client-id"
	ClientSecretFlag  string = "client-secret"
	APIEndpointFlag   string = "api-endpoint"
	TokenEndpointFlag string = "token-endpoint"
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

	cmd.Flags().StringP(ClientIdFlag,
		"c",
		"",
		"the client id to call the API")
	cmd.MarkFlagRequired(ClientIdFlag)

	cmd.Flags().StringP(ClientSecretFlag,
		"s",
		"",
		"the client secret to call the API")
	cmd.MarkFlagRequired(ClientSecretFlag)

	cmd.Flags().StringP(APIEndpointFlag,
		"a",
		"",
		"the API endpoint")
	cmd.MarkFlagRequired(APIEndpointFlag)

	cmd.Flags().StringP(TokenEndpointFlag,
		"t",
		"",
		"the endpoint to get authentication tokens")
	cmd.MarkFlagRequired(TokenEndpointFlag)

	return cmd
}

// execute implements all the logic associated with this command.
func execute(cmd *cobra.Command, args []string) error {
	clientId, _ := cmd.Flags().GetString(ClientIdFlag)
	viper.Set(ClientIdFlag, clientId)

	clientSecret, _ := cmd.Flags().GetString(ClientSecretFlag)
	viper.Set(ClientSecretFlag, clientSecret)

	apiEndpoint, _ := cmd.Flags().GetString(APIEndpointFlag)
	viper.Set(APIEndpointFlag, apiEndpoint)

	tokenEndpoint, _ := cmd.Flags().GetString(TokenEndpointFlag)
	viper.Set(TokenEndpointFlag, tokenEndpoint)

	return viper.WriteConfig()
}

// initConfig reads in config file and ENV variables if set
func initConfig() {

	// find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// search config in home directory with name
	// ".learning-go-cli" (without extension).
	ext := "yaml"
	name := ".learning-go-cli"
	viper.AddConfigPath(home)
	viper.SetConfigType(ext)
	viper.SetConfigName(name)

	// creates config file if it does not exist
	_, err = createConfigFile(home, name, ext)
	cobra.CheckErr(err)

	// if a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		cobra.CheckErr(err)
	}
}

// createConfigFile creates the config file if it does not exist
func createConfigFile(home string, name string, ext string) (string, error) {
	fileName := fmt.Sprintf("%s/%s.%s", home, name, ext)
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		file, err := os.Create(fileName)
		if err != nil {
			defer file.Close()
			return fileName, nil
		}
		return fileName, err
	}
	return fileName, nil
}

// addCommandWithConfigPreCheck adds a command to the rootCmd configuring a
// PreRunE function to ensure the configure command is executed before
// any other command
func addCommandWithConfigPreCheck(cmd *cobra.Command) {
	cmd.PreRunE = configPreCheck
	rootCmd.AddCommand(cmd)
}

// configPreCheck verifies if the base configuration is set
func configPreCheck(cmd *cobra.Command, args []string) error {
	validConfig := viper.InConfig(ClientIdFlag) &&
		viper.InConfig(ClientSecretFlag) &&
		viper.InConfig(APIEndpointFlag) &&
		viper.InConfig(TokenEndpointFlag)

	fmt.Printf("flag = %s, %s\n", viper.Get(TokenEndpointFlag), viper.ConfigFileUsed())

	if !validConfig {
		return fmt.Errorf(
			"invalid CLI configuration: " +
				"please run `learning-go-api configure`")
	}

	return nil
}
