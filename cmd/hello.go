package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/maxvoronov/go-release/internal/greeter"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Show information about version, commit and build time",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(greeter.Hello(args[0]))
	},
}
