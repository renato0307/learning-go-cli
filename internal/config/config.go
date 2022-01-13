package config

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

const (
	ClientIdFlag      string = "client-id"
	ClientSecretFlag  string = "client-secret"
	APIEndpointFlag   string = "api-endpoint"
	TokenEndpointFlag string = "token-endpoint"
)

// GetString returns a configuration string
func GetString(key string) string {
	return viper.GetString(key)
}

// Set defines a configuration value
func Set(key string, value interface{}) {
	viper.Set(key, value)
}

// WriteAuthenticationConfig persists the authentication configuration
func WriteAuthenticationConfig(
	clientId,
	clientSecret,
	apiEndpoint,
	tokenEndpoint string) error {

	Set(ClientSecretFlag, clientSecret)
	Set(ClientIdFlag, clientId)
	Set(APIEndpointFlag, apiEndpoint)
	Set(TokenEndpointFlag, tokenEndpoint)

	return viper.WriteConfig()
}

// initConfig reads in config file and ENV variables if set
func InitConfig() {

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

// configPreCheck verifies if the base configuration is set
func ConfigPreCheck(cmd *cobra.Command, args []string) error {
	validConfig := viper.InConfig(ClientIdFlag) &&
		viper.InConfig(ClientSecretFlag) &&
		viper.InConfig(APIEndpointFlag) &&
		viper.InConfig(TokenEndpointFlag)

	fmt.Printf("flag = %s, %s\n",
		viper.Get(TokenEndpointFlag),
		viper.ConfigFileUsed())

	if !validConfig {
		return fmt.Errorf(
			"invalid CLI configuration: " +
				"please run `learning-go-api configure`")
	}

	return nil
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

// CreateFakeConfigFile configures viper to write to a temporary file
func CreateFakeConfigFile(t *testing.T) string {
	home, _ := os.UserHomeDir()
	ext := "yaml"
	name := fmt.Sprintf(".learning-go-cli-test-%d", rand.Uint64())
	fileName, err := createConfigFile(home, name, ext)
	if err != nil {
		assert.FailNow(t, "error creating config file")
	}

	viper.Reset()
	viper.AddConfigPath(home)
	viper.SetConfigType(ext)
	viper.SetConfigName(name)

	Set(APIEndpointFlag, "fake_endpoint")
	Set(TokenEndpointFlag, "fake_endpoint")
	Set(ClientIdFlag, "fake_client_id")
	Set(ClientSecretFlag, "fake_client_secret")
	viper.WriteConfig()

	return fileName
}
