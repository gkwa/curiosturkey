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

	allPaths, err := collectAllPaths(expandedPath)
	if err != nil {
		return nil, err
	}

	repoInfos, err := processGitRepos(allPaths)
	if err != nil {
		return nil, err
	}

	logger.V(1).Info("Finished collecting repository information", "count", len(repoInfos))

	return repoInfos, nil
}

func collectAllPaths(rootDir string) ([]string, error) {
	var allPaths []string

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		allPaths = append(allPaths, path)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error walking the path %s: %v", rootDir, err)
	}

	return allPaths, nil
}

func processGitRepos(paths []string) ([]RepoInfo, error) {
	var repoInfos []RepoInfo

	for _, path := range paths {
		info, err := os.Stat(path)
		if err != nil {
			continue
		}

		if isGitRepo(info) {
			repoInfo, err := processRepo(path)
			if err != nil {
				continue
			}
			repoInfos = append(repoInfos, repoInfo)
		}
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
