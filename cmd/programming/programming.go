package programming

import (
	"fmt"

	"github.com/renato0307/learning-go-cli/internal/config"
	"github.com/renato0307/learning-go-cli/internal/iostreams"
	"github.com/spf13/cobra"
)

// NewProgrammingCmd represents the programming command
func NewProgrammingCmd(iostreams *iostreams.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "programming",
		Short: "Programming tools",
		Long:  `Provides several programming tools like uuid generation, etc.`,
		RunE:  executeProgramming(),
	}

	config.AddCommandWithConfigPreCheck(cmd, NewProgrammingUuidCmd(iostreams))

	return cmd
}

// executeProgramming implements all the logic associated with this command.
// In this case as it is an aggregation command will return an error
func executeProgramming() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("must specify a subcommand")
	}
}
