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

type runner struct {
	userDir string
}

func Run(config map[string]interface{}) error {
	var userDir string
	if runtime.GOOS == "windows" {
		userDir = os.Getenv("USERPROFILE")
	} else {
		userDir = os.Getenv("HOME")
	}

	return (&runner{
		userDir: userDir,
	}).Run(config)
}

func (r *runner) Run(config map[string]interface{}) error {
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

	depthConfig, _ := config["depth"]
	depth := 0
	if depthConfig != nil {
		depth = depthConfig.(int)
	}

	needsSSH := !strings.HasPrefix(repoURL, "http:") && !strings.HasPrefix(repoURL, "https:")
	var auth transport.AuthMethod
	if needsSSH {
		var err error
		auth, err = r.getSSHAuth(repoURL)
		if err != nil {
			return err
		}
	}

	repo, err := git.PlainClone(dest, false, &git.CloneOptions{
		Auth:  auth,
		URL:   repoURL,
		Depth: depth,
	})
	if err != nil {
		if !errors.Is(err, git.ErrRepositoryAlreadyExists) {
			return fmt.Errorf("Failed cloning %s to %s: %s", repoURL, dest, err)
		}

		repo, err = git.PlainOpen(dest)
		if err != nil {
			return fmt.Errorf("Failed accessing git worktree at %s: %s", dest, err)
		}
		err = repo.Fetch(&git.FetchOptions{
			Depth: depth,
		})
		if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
			if depth == 0 {
				return fmt.Errorf("Fetch failed: %s", err)
			} else {
				return fmt.Errorf("Fetch failed with {depth: %d}: %s", depth, err)
			}
		}
	}

	wt, err := repo.Worktree()
	if err != nil {
		return fmt.Errorf("Failed returning Worktree: %s", err)
	}

	err = wt.Checkout(&git.CheckoutOptions{
		Hash: plumbing.NewHash(version),
	})
	if err != nil {
		return fmt.Errorf("Failed checking out %s: %s", version, err)
	}

	return nil
}

func (r *runner) getSSHPublicKeys(url string) (*gossh.PublicKeys, error) {

	pk := filepath.Join(r.userDir, ".ssh", "id_rsa")

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

func (r *runner) getSSHAuth(url string) (transport.AuthMethod, error) {
	return r.getSSHPublicKeys(url)
}
