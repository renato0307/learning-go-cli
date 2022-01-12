package cmd

import (
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/renato0307/learning-go-cli/internal/config"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func createFakeConfigFile(t *testing.T) {
	home, _ := os.UserHomeDir()
	ext := "yaml"
	name := fmt.Sprintf(".learning-go-cli-test-%d", rand.Uint64())
	fileName, err := createConfigFile(home, name, ext)
	if err != nil {
		assert.FailNow(t, "error creating config file")
	}
	defer os.Remove(fileName)
	viper.Reset()
	viper.AddConfigPath(home)
	viper.SetConfigType(ext)
	viper.SetConfigName(name)
}

func TestNewConfigureCommand(t *testing.T) {

	// act
	cmd := NewConfigureCommand()

	// assert
	assert.Equal(t, "configure", cmd.Use)
	assert.NotEmpty(t, cmd.Short, "Short description cannot be empty")
	assert.NotEmpty(t, cmd.Long, "Long description cannot be empty")
	assert.NotNil(t, cmd.RunE, "The RunE function must be defined")
	assert.NotNil(t, cmd.Flags().Lookup(config.ClientIdFlag))
	assert.NotNil(t, cmd.Flags().Lookup(config.ClientSecretFlag))
	assert.NotNil(t, cmd.Flags().Lookup(config.APIEndpointFlag))
	assert.NotNil(t, cmd.Flags().Lookup(config.TokenEndpointFlag))
}

func TestExecute(t *testing.T) {

	// arrange
	cmd := NewConfigureCommand()
	createFakeConfigFile(t)

	// act
	cmd.SetArgs([]string{
		"-c", "fake-c",
		"-s", "fake-s",
		"-a", "fake-a",
		"-t", "fake-t",
	})
	err := cmd.Execute()

	// assert
	assert.NoError(t, err)
	assert.Equal(t, "fake-c", viper.GetViper().Get(config.ClientIdFlag))
	assert.Equal(t, "fake-s", viper.GetViper().Get(config.ClientSecretFlag))
	assert.Equal(t, "fake-a", viper.GetViper().Get(config.APIEndpointFlag))
	assert.Equal(t, "fake-t", viper.GetViper().Get(config.TokenEndpointFlag))
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

func TestAddCommandWithConfigPreCheck(t *testing.T) {

	// arrange
	cmd := NewConfigureCommand()

	// act
	addCommandWithConfigPreCheck(cmd)

	// assert
	assert.NotNil(t, cmd.PreRunE)
	found := false
	for _, c := range rootCmd.Commands() {
		found = (c == cmd)
		if found {
			break
		}
	}
	assert.True(t, found, "command was not added to the rootCmd")
}

func TestConfigPreCheckReturnsErrorIfMissingConfigs(t *testing.T) {

	// arrange
	cmd := NewConfigureCommand()
	createFakeConfigFile(t)
	viper.ReadInConfig()

	// act
	err := configPreCheck(cmd, []string{})

	// assert
	print(err)
	assert.Error(t, err)
}

func TestConfigPreCheckReturnsNoErrorIfConfigsFound(t *testing.T) {

	// arrange
	cmd := NewConfigureCommand()
	createFakeConfigFile(t)
	cmd.SetArgs([]string{
		"-c", "fake-c",
		"-s", "fake-s",
		"-a", "fake-a",
		"-t", "fake-t",
	})
	cmd.Execute()
	viper.ReadInConfig()

	// act
	err := configPreCheck(cmd, []string{})

	// assert
	assert.NoError(t, err)
}
