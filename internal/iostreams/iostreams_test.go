package iostreams

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintOutput(t *testing.T) {
	// arrange
	buffer := &bytes.Buffer{}
	iostreams := IOStreams{Out: buffer}
	s := "my test string"

	// act
	iostreams.Fprint(s)

	// assert
	assert.Equal(t, s, buffer.String())
}
