package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
