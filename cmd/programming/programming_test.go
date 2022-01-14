package programming

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProgrammingCmd(t *testing.T) {
	// act
	cmd := NewProgrammingCmd(nil)

	// assert
	assert.Equal(t, "programming", cmd.Use)
	assert.NotEmpty(t, cmd.Short, "Short description cannot be empty")
	assert.NotEmpty(t, cmd.Long, "Long description cannot be empty")
	assert.NotNil(t, cmd.RunE, "The RunE function must be defined")
}

func TestExecute(t *testing.T) {
	// arrange
	cmd := NewProgrammingCmd(nil)

	// act
	err := cmd.Execute()

	// assert
	assert.Error(t, err)
}
