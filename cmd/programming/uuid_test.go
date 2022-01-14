package programming

import (
	"bytes"
	"fmt"
	"net/http"
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
	assert.NotNil(t, cmd.RunE, "The RunE function must be defined")
	assert.NotNil(t, cmd.Flags().Lookup(NoHyphensFlag))
}

func TestExecuteProgrammingUuid(t *testing.T) {

	uuid := "da308fbd-cba9-485a-b4c1-6677aaa732a4"
	uuidNoHyphens := "da308fbdcba9485ab4c16677aaa732a4"

	testCases := []struct {
		ApiStatusCode     int
		ApiResponse       string
		ApiExpectedParams []string
		OutputContains    string
		Args              []string
		ErrorNil          bool
		Purpose           string
	}{
		{
			ApiStatusCode:  http.StatusOK,
			ApiResponse:    fmt.Sprintf("{\"uuid\": \"%s\"}", uuid),
			Args:           []string{},
			OutputContains: uuid,
			ErrorNil:       true,
			Purpose:        "success case",
		},
		{
			ApiStatusCode:     http.StatusOK,
			ApiResponse:       fmt.Sprintf("{\"uuid\": \"%s\"}", uuidNoHyphens),
			ApiExpectedParams: []string{"no-hyphens"},
			OutputContains:    uuidNoHyphens,
			Args:              []string{fmt.Sprintf("--%s", NoHyphensFlag)},
			ErrorNil:          true,
			Purpose:           "success case with no hyphens",
		},
		{
			ApiStatusCode: http.StatusBadRequest,
			ApiResponse:   "{\"message\": \"request is malformed\"}",
			Args:          []string{},
			ErrorNil:      false,
			Purpose:       "api returns error",
		},
		{
			ApiStatusCode: http.StatusOK,
			ApiResponse:   "something that is not a valid json",
			Args:          []string{},
			ErrorNil:      false,
			Purpose:       "error on invalid json",
		},
	}

	for _, tc := range testCases {
		// arrange
		buffer := &bytes.Buffer{}
		iostreams := &iostreams.IOStreams{Out: buffer}
		cmd := NewProgrammingUuidCmd(iostreams)

		tokenSrv := testhelpers.NewAuthTestServer()
		defer tokenSrv.Close()
		config.Set(config.TokenEndpointFlag, tokenSrv.URL)

		apiSrv := testhelpers.NewAPITestServer(
			tc.ApiResponse,
			tc.ApiExpectedParams,
			tc.ApiStatusCode)
		defer apiSrv.Close()
		config.Set(config.APIEndpointFlag, apiSrv.URL)

		// act
		cmd.SetArgs(tc.Args)
		err := cmd.Execute()

		// assert
		if tc.ErrorNil {
			assert.NoError(t, err, "error found for "+tc.Purpose)
			assert.Contains(t,
				buffer.String(),
				tc.OutputContains,
				"output is for right for "+tc.Purpose)
		} else {
			assert.Error(t, err, "error not found for "+tc.Purpose)
		}
	}
}
