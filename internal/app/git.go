package app

import (
	"os/exec"
	"strings"
)

func CurrentGitBranch(dir string) (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = dir
	buf, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(buf)), nil
}
