package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/gkwa/curiosturkey/core"
	"github.com/spf13/cobra"
)

var newerThanCmd = &cobra.Command{
	Use:   "newerthan [timespan] [path]",
	Short: "Order repositories newer than the specified timespan",
	Long:  `This command orders the repositories in the given path that are newer than the specified timespan.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		timespan, err := core.ParseTimespan(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing timespan: %v\n", err)
			os.Exit(1)
		}

		repoInfos, err := core.OrderReposByCommitDate(cmd.Context(), args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		filteredRepoInfos := filterReposByTimespan(repoInfos, timespan)
		printRepoInfo(filteredRepoInfos)
	},
}

func filterReposByTimespan(repoInfos []core.RepoInfo, timespan time.Duration) []core.RepoInfo {
	now := time.Now()
	var filtered []core.RepoInfo
	for _, info := range repoInfos {
		if now.Sub(info.LatestDate) <= timespan {
			filtered = append(filtered, info)
		}
	}
	return filtered
}

func printRepoInfo(repoInfos []core.RepoInfo) {
	now := time.Now()
	for _, info := range repoInfos {
		relTime := core.FormatUserFriendlyDuration(now.Sub(info.LatestDate))
		fmt.Printf("%s: %s\n", relTime, info.Path)
	}
}

func init() {
	rootCmd.AddCommand(newerThanCmd)
}
