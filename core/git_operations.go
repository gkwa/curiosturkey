package core

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
)

func getLatestCommitDate(repoPath string) (time.Time, error) {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return time.Time{}, fmt.Errorf("error opening repository: %v", err)
	}

	ref, err := repo.Head()
	if err != nil {
		return time.Time{}, fmt.Errorf("error getting HEAD: %v", err)
	}

	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return time.Time{}, fmt.Errorf("error getting commit: %v", err)
	}

	return commit.Author.When, nil
}
