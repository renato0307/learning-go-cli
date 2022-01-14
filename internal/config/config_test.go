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

func TestGetString(t *testing.T) {
	// arrange
	key := "someconfig"
	value := "somevalue"
	viper.Set(key, value)

	// act
	returnValue := GetString(key)

	// assert
	assert.Equal(t, value, returnValue)
}

func TestSet(t *testing.T) {
	// arrange
	key := "someconfig"
	value := "somevalue"

	// act
	Set(key, value)

	// assert
	returnValue := viper.Get(key)
	assert.Equal(t, value, returnValue)
}

func TestCreateConfigFile(t *testing.T) {
	// arrange
	home, _ := os.UserHomeDir()
	ext := "yaml"
	name := fmt.Sprintf(".learning-go-cli-test-%d", rand.Uint64())

	// act
	fileName, err := createConfigFile(home, name, ext)

	// assert
	if err != nil {
		assert.Fail(t, "error creating config file")
	}
	defer os.Remove(fileName)
	assert.Equal(t, fmt.Sprintf("%s/%s.%s", home, name, ext), fileName)
}

func TestConfigPreCheckReturnsErrorIfMissingConfigs(t *testing.T) {
	// arrange
	fileName := CreateFakeConfigFile(t)
	defer os.Remove(fileName)

	// act
	err := ConfigPreCheck(&cobra.Command{}, []string{})

	// assert
	print(err)
	assert.Error(t, err)
}

func TestConfigPreCheckReturnsNoErrorIfConfigsFound(t *testing.T) {
	// arrange
	fileName := CreateFakeConfigFile(t)
	defer os.Remove(fileName)
	viper.ReadInConfig()

	// act
	err := ConfigPreCheck(&cobra.Command{}, []string{})

	// assert
	assert.NoError(t, err)
}

func TestWriteAuthenticationConfig(t *testing.T) {
	// arrange
	fileName := CreateFakeConfigFile(t)
	defer os.Remove(fileName)

	apiEndpoint := "fake_endpoint_2"
	tokenEndpoint := "fake_endpoint_2"
	clientId := "fake_client_id_2"
	clientSecret := "fake_client_secret_2"

	// act
	WriteAuthenticationConfig(clientId,
		clientSecret,
		apiEndpoint,
		tokenEndpoint)

	// assert
	assert.Equal(t, GetString(ClientIdFlag), clientId)
	assert.Equal(t, GetString(ClientSecretFlag), clientSecret)
	assert.Equal(t, GetString(TokenEndpointFlag), tokenEndpoint)
	assert.Equal(t, GetString(APIEndpointFlag), apiEndpoint)
}

func TestInitConfig(t *testing.T) {
	// act
	InitConfig()

	// assert
	assert.NotEmpty(t, viper.ConfigFileUsed())
	assert.Contains(t, viper.ConfigFileUsed(), "learning-go-cli")
}

func TestAddCommandWithConfigPreCheck(t *testing.T) {
	// arrange
	cmd := &cobra.Command{}
	parentCmd := &cobra.Command{}

	// act
	AddCommandWithConfigPreCheck(parentCmd, cmd)

	// assert
	assert.NotNil(t, cmd.PreRunE)
	assert.Contains(t, parentCmd.Commands(), cmd)
}
