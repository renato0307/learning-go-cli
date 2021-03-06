package cmd

import (
	"os"

	"github.com/renato0307/learning-go-cli/cmd/programming"
	"github.com/renato0307/learning-go-cli/internal/config"
	"github.com/renato0307/learning-go-cli/internal/iostreams"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "learning-go-cli",
	Short: "CLI for the learning-go-api",
	Long: `The learning-go-api provides with utility functions like UUID
generation, a currency converter, a JWT debugger, etc.`,
	Version: "0.0.1",
}

// Execute adds all child commands to the root command and sets flags
// appropriately. This is called by main.main(). It only needs to happen once
// to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(config.InitConfig)

	iostreams := &iostreams.IOStreams{Out: os.Stdout}

	rootCmd.AddCommand(NewConfigureCommand(iostreams))

	programmingCmd := programming.NewProgrammingCmd(iostreams)
	config.AddCommandWithConfigPreCheck(rootCmd, programmingCmd)
}
