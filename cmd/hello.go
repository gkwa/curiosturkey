package cmd

import (
	"fmt"
	"os"

	"github.com/gkwa/curiosturkey/core"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello [path]",
	Short: "Order repositories by commit date",
	Long:  `This command orders the repositories in the given path by their most recent commit date.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logger := LoggerFrom(cmd.Context())
		logger.Info("Running hello command")
		err := core.OrderReposByCommitDate(cmd.Context(), args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
