package cmd

import (
	"fmt"
	"os"

	"github.com/renato0307/learning-go-cli/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	viper.Set(config.ClientIdFlag, clientId)

	clientSecret, _ := cmd.Flags().GetString(config.ClientSecretFlag)
	viper.Set(config.ClientSecretFlag, clientSecret)

	apiEndpoint, _ := cmd.Flags().GetString(config.APIEndpointFlag)
	viper.Set(config.APIEndpointFlag, apiEndpoint)

	tokenEndpoint, _ := cmd.Flags().GetString(config.TokenEndpointFlag)
	viper.Set(config.TokenEndpointFlag, tokenEndpoint)

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
	validConfig := viper.InConfig(config.ClientIdFlag) &&
		viper.InConfig(config.ClientSecretFlag) &&
		viper.InConfig(config.APIEndpointFlag) &&
		viper.InConfig(config.TokenEndpointFlag)

	fmt.Printf("flag = %s, %s\n",
		viper.Get(config.TokenEndpointFlag),
		viper.ConfigFileUsed())

	if !validConfig {
		return fmt.Errorf(
			"invalid CLI configuration: " +
				"please run `learning-go-api configure`")
	}

	return nil
}
