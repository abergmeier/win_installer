package git

import (
	"errors"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func Run(config map[string]interface{}) error {
	repoConfig, ok := config["repo"]
	if !ok {
		return errors.New("Missing repo config")
	}
	if repoConfig == nil {
		return errors.New("Missing value in repo config")
	}
	repoURL := repoConfig.(string)
	dest := config["dest"].(string)
	version := config["version"].(string)

	repo, err := git.PlainClone(dest, false, &git.CloneOptions{
		URL: repoURL,
	})
	if err != nil {
		if !errors.Is(err, git.ErrRepositoryAlreadyExists) {
			return err
		}

		repo, err = git.PlainOpen(dest)
		if err != nil {
			return err
		}
		err = repo.Fetch(&git.FetchOptions{})
		if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
			return err
		}
	}

	wt, err := repo.Worktree()
	if err != nil {
		return err
	}

	return wt.Checkout(&git.CheckoutOptions{
		Hash: plumbing.NewHash(version),
	})
}
