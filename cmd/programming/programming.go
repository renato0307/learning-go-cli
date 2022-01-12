package programming

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewProgrammingCmd represents the programming command
func NewProgrammingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "programming",
		Short: "Programming tools",
		Long:  `Provides several programming tools like uuid generation, etc.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return execute(cmd, args)
		},
	}

	cmd.AddCommand(NewProgrammingUuidCmd())

	return cmd
}

// execute implements all the logic associated with this command.
// In this case as it is an aggregation command will return an error
func execute(cmd *cobra.Command, args []string) error {
	return fmt.Errorf("must specify a subcommand")
}
