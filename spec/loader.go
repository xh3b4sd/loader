package spec

import (
	"github.com/spf13/cobra"
)

// Loader represents the business logic implementation usable in command line
// tools.
type Loader interface {
	// Boot initializes and executes the command line tool.
	Boot()

	// ExecGenerateCmd executes the generate command.
	ExecGenerateCmd(cmd *cobra.Command, args []string)

	// ExecVersionCmd executes the version command.
	ExecVersionCmd(cmd *cobra.Command, args []string)

	// InitGenerateCmd initializes the generate command.
	InitGenerateCmd() *cobra.Command

	// InitVersionCmd initializes the version command.
	InitVersionCmd() *cobra.Command
}
