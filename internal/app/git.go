package app

import (
	"os/exec"
	"strings"
)

func (l *LogRunner) CurrentGitBranch(dir string) (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = dir
	buf, err := cmd.Output()
	if err != nil {
		l.Fatal("failed getting the git branch: %v", err)
		return "", err
	}

	return strings.TrimSpace(string(buf)), nil
}
