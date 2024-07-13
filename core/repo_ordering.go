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

	allPaths, err := collectAllPaths(ctx, expandedPath)
	if err != nil {
		return nil, err
	}

	repoInfos, err := processGitRepos(ctx, allPaths)
	if err != nil {
		return nil, err
	}

	logger.V(1).Info("Finished collecting repository information", "count", len(repoInfos))

	return repoInfos, nil
}

func collectAllPaths(ctx context.Context, rootDir string) ([]string, error) {
	logger := logr.FromContextOrDiscard(ctx)
	var allPaths []string

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		logger.V(1).Info("Found path", "path", path)
		allPaths = append(allPaths, path)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error walking the path %s: %v", rootDir, err)
	}

	logger.V(1).Info("Collected all paths", "count", len(allPaths))
	return allPaths, nil
}

func processGitRepos(ctx context.Context, paths []string) ([]RepoInfo, error) {
	logger := logr.FromContextOrDiscard(ctx)
	var repoInfos []RepoInfo

	for _, path := range paths {
		info, err := os.Stat(path)
		if err != nil {
			continue
		}

		if isGitRepo(info) {
			absPath, err := filepath.Abs(path)
			if err != nil {
				logger.Error(err, "Failed to get absolute path", "path", path)
				continue
			}
			logger.V(1).Info("Discovered Git repository", "path", absPath)

			repoInfo, err := processRepo(path)
			if err != nil {
				logger.Error(err, "Failed to process repository", "path", absPath)
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
