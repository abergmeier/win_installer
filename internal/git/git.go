package git

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport"
	gossh "github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"golang.org/x/crypto/ssh"
)

var (
	sshRegex = regexp.MustCompile(`(?:ssh://){0,1}(user){0,1}(?:@){0,1}.+/.+`)
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

	needsSSH := !strings.HasPrefix(repoURL, "http:") && !strings.HasPrefix(repoURL, "https:")
	var auth transport.AuthMethod
	if needsSSH {
		var err error
		auth, err = getSSHAuth(repoURL)
		if err != nil {
			return err
		}
	}

	repo, err := git.PlainClone(dest, false, &git.CloneOptions{
		Auth: auth,
		URL:  repoURL,
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

func getSSHPublicKeys(url string) (*gossh.PublicKeys, error) {

	var pk string
	if runtime.GOOS == "windows" {
		pk = filepath.Join(os.Getenv("USERPROFILE"), ".ssh", "id_rsa")
	} else {
		pk = filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa")
	}

	matches := sshRegex.FindStringSubmatch(url)
	if len(matches) == 0 {
		return nil, fmt.Errorf("Could not handle URL via regex: %s", url)
	}

	user := matches[1]

	sshKey, err := ioutil.ReadFile(pk)
	if err != nil {
		return nil, err
	}
	signer, err := ssh.ParsePrivateKey(sshKey)
	if err != nil {
		return nil, err
	}
	return &gossh.PublicKeys{User: user, Signer: signer}, nil
}

func getSSHAuth(url string) (transport.AuthMethod, error) {
	return getSSHPublicKeys(url)
}
