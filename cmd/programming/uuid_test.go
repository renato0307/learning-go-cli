package programming

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/renato0307/learning-go-cli/internal/config"
	"github.com/renato0307/learning-go-cli/internal/iostreams"
	"github.com/renato0307/learning-go-cli/internal/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestNewProgrammingUuidCmd(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	iostreams := &iostreams.IOStreams{Out: buffer}

	// act
	cmd := NewProgrammingUuidCmd(iostreams)

	// assert
	assert.Equal(t, "uuid", cmd.Use)
	assert.NotEmpty(t, cmd.Short, "Short description cannot be empty")
	assert.NotEmpty(t, cmd.Long, "Long description cannot be empty")
	assert.NotNil(t, cmd.Run, "The Run function must be defined")
	assert.NotNil(t, cmd.Flags().Lookup(NoHiphensFlag))
}

func TestExecuteProgrammingUuid(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	iostreams := &iostreams.IOStreams{Out: buffer}
	cmd := NewProgrammingUuidCmd(iostreams)

	tokenSrv := testhelpers.NewAuthTestServer()
	config.Set(config.TokenEndpointFlag, tokenSrv.URL)

	uuid := "da308fbd-cba9-485a-b4c1-6677aaa732a4"
	apiResponse := fmt.Sprintf("{\"uuid\": \"%s\"}", uuid)
	apiSrv := testhelpers.NewAPITestServer(apiResponse)
	config.Set(config.APIEndpointFlag, apiSrv.URL)

	// act
	executeProgrammingUuid(cmd, []string{}, iostreams)

	// assert
	assert.Contains(t, buffer.String(), uuid)
}
