package iostreams

import (
	"fmt"
	"io"
)

// IOStreams represents the structures needed for input/output in commands
// Currently only supports output.
type IOStreams struct {
	Out io.Writer
}

// PrintOutput knows how to print using an IOStreams struct
func (iostreams *IOStreams) Fprint(v interface{}) (n int, err error) {
	return fmt.Fprint(iostreams.Out, v)
}
