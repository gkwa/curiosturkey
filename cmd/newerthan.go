package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/gkwa/curiosturkey/core"
	"github.com/spf13/cobra"
)

var (
	hideAge bool
	age     string
)

var newerThanCmd = &cobra.Command{
	Use:   "newerthan <path> [<path>...] --age=<timespan>",
	Short: "Order repositories newer than the specified timespan",
	Long:  `This command orders the repositories in the given paths that are newer than the specified timespan.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if age == "" {
			return fmt.Errorf("--age flag is required")
		}

		timespan, err := core.ParseTimespan(age)
		if err != nil {
			return fmt.Errorf("error parsing timespan: %v", err)
		}

		var allRepoInfos []core.RepoInfo
		for _, path := range args {
			repoInfos, err := core.OrderReposByCommitDate(cmd.Context(), path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error processing path %s: %v\n", path, err)
				continue
			}
			allRepoInfos = append(allRepoInfos, repoInfos...)
		}

		filteredRepoInfos := filterReposByTimespan(allRepoInfos, timespan)
		core.SortRepoInfos(filteredRepoInfos)
		printRepoInfo(filteredRepoInfos)

		return nil
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
		if hideAge {
			fmt.Printf("%s\n", info.Path)
		} else {
			relTime := core.FormatUserFriendlyDuration(now.Sub(info.LatestDate))
			fmt.Printf("%s: %s\n", relTime, info.Path)
		}
	}
}

func init() {
	rootCmd.AddCommand(newerThanCmd)
	newerThanCmd.Flags().BoolVar(&hideAge, "hide-age", false, "Hide the age of the repositories")
	newerThanCmd.Flags().StringVar(&age, "age", "", "Age threshold for repositories (e.g., 1y2M3w4d5h6m)")
	if err := newerThanCmd.MarkFlagRequired("age"); err != nil {
		fmt.Fprintf(os.Stderr, "Error marking 'age' flag as required: %v\n", err)
		os.Exit(1)
	}
}
