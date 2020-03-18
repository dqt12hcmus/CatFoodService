package cmd

import "github.com/spf13/cobra"

const orderServiceName = "order"

var orderCmd = &cobra.Command{
	Use:   orderServiceName,
	Short: "Start order service",
	Long:  "Start order service to receive order from customers",
	RunE: func(cmd *cobra.Command, args []string) error {
		cleanUpLog := log.SetupLogging(orderServiceName)
		defer cleanUpLog()

		return orderservice.Run()
	},
}
