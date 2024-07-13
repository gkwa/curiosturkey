package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/gkwa/curiosturkey/core"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello [path]",
	Short: "Order repositories by commit date",
	Long:  `This command orders the repositories in the given path by their most recent commit date.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		repoInfos, err := core.OrderReposByCommitDate(cmd.Context(), args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		printRepoInfo(repoInfos)
	},
}

func printRepoInfo(repoInfos []core.RepoInfo) {
	now := time.Now()
	for _, info := range repoInfos {
		relTime := core.FormatUserFriendlyDuration(now.Sub(info.LatestDate))
		fmt.Printf("%s: %s\n", relTime, info.Path)
	}
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
