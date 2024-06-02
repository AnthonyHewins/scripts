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

func (l *LogRunner) GitBranches() ([]string, error) {
	cmd := exec.Command("git", "branch")
	buf, err := cmd.Output()
	if err != nil {
		l.Fatal("failed getting list of git branches")
		return nil, err
	}

	branches := strings.Split(string(buf), "\n")
	for i, v := range branches {
		branches[i] = strings.TrimSpace(strings.TrimPrefix(v, "* "))
	}

	return branches, nil
}
