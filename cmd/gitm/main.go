package main

import (
	"os"

	"github.com/AnthonyHewins/scripts/internal/app"
)

func main() {
	dir := "."
	l := app.NewLogRunner(dir, os.Stdout, os.Stderr)

	branch, err := app.CurrentGitBranch(dir)
	if err != nil {
		l.Fatal("Failed fetching git branch name: %s", err.Error())
	}

	if branch != "master" {
		l.Run("git", "checkout", "master")
		l.Run("git", "pull", "origin", "master")
		l.Run("git", "checkout", branch)
	}

	l.Run("git", "pull", "origin", "master")
}
