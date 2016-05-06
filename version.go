package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (l *loader) InitVersionCmd() *cobra.Command {
	newCmd := &cobra.Command{
		Use:   "version",
		Short: "Show current version of the binary.",
		Long:  "Show current version of the binary.",
		Run:   l.ExecVersionCmd,
	}

	return newCmd
}

func (l *loader) ExecVersionCmd(cmd *cobra.Command, args []string) {
	fmt.Printf("%s\n", l.Version)
}
