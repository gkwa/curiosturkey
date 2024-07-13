package core

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/go-logr/logr"
	"github.com/mitchellh/go-homedir"
)

func OrderReposByCommitDate(ctx context.Context, rootDir string) ([]RepoInfo, error) {
	logger := logr.FromContextOrDiscard(ctx)

	expandedPath, err := homedir.Expand(rootDir)
	if err != nil {
		return nil, fmt.Errorf("error expanding path: %v", err)
	}

	repoInfos, err := CollectReposInfo(expandedPath)
	if err != nil {
		return nil, err
	}

	logger.V(1).Info("Finished collecting repository information", "count", len(repoInfos))

	return repoInfos, nil
}

func CollectReposInfo(rootDir string) ([]RepoInfo, error) {
	var repoInfos []RepoInfo

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if isGitRepo(info) {
			repoInfo, err := processRepo(path)
			if err != nil {
				return err
			}
			repoInfos = append(repoInfos, repoInfo)
			return filepath.SkipDir
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error walking the path %s: %v", rootDir, err)
	}

	return repoInfos, nil
}

func isGitRepo(info os.FileInfo) bool {
	return info.IsDir() && info.Name() == ".git"
}

func processRepo(path string) (RepoInfo, error) {
	repoPath := filepath.Dir(path)
	latestDate, err := GetLatestCommitDate(repoPath)
	if err != nil {
		return RepoInfo{}, fmt.Errorf("error processing repository %s: %v", repoPath, err)
	}

	return RepoInfo{
		Path:       repoPath,
		LatestDate: latestDate,
	}, nil
}

func SortRepoInfos(repoInfos []RepoInfo) {
	sort.Slice(repoInfos, func(i, j int) bool {
		return repoInfos[i].LatestDate.Before(repoInfos[j].LatestDate)
	})
}
