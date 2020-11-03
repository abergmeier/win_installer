package git

import (
	"testing"
)

func TestClone(t *testing.T) {

	err := Run(map[string]interface{}{
		"repo":    "git@github.com:go-git/go-git.git",
		"dest":    "/tmp/go-git-checkout",
		"version": "d525a514057f97bc2b183e2c67f542dd6f0ac0aa",
	})
	if err != nil {
		t.Fatal(err)
	}

}

func TestSSHAuth(t *testing.T) {
	pk, err := getSSHPublicKeys("ssh://user@server/project.git")
	if err != nil {
		t.Fatal(err)
	}

	if pk.User != "user" {
		t.Fatalf("Unexpected user: %s", pk.User)
	}
}
