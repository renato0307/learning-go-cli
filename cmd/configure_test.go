package cmd

import (
	"bytes"
	"os"
	"testing"

	"github.com/renato0307/learning-go-cli/internal/config"
	"github.com/renato0307/learning-go-cli/internal/iostreams"
	"github.com/stretchr/testify/assert"
)

func TestNewConfigureCommand(t *testing.T) {
	// act
	cmd := NewConfigureCommand(nil)

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

func TestExecuteConfigure(t *testing.T) {
	// arrange
	// arrange
	buffer := &bytes.Buffer{}
	iostreams := &iostreams.IOStreams{Out: buffer}
	cmd := NewConfigureCommand(iostreams)

	fileName := config.CreateFakeConfigFile(t)
	defer os.Remove(fileName)

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
	assert.Equal(t, "fake-c", config.GetString(config.ClientIdFlag))
	assert.Equal(t, "fake-s", config.GetString(config.ClientSecretFlag))
	assert.Equal(t, "fake-a", config.GetString(config.APIEndpointFlag))
	assert.Equal(t, "fake-t", config.GetString(config.TokenEndpointFlag))
	assert.Equal(t, "configuration updated!", buffer.String())
}
