package cmd

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "cat-shop",
}

func init() {
	Cmd.AddCommand(schedulerCmd)
}
