package main

import (
	"fmt"
	"os"

	"github.com/AnthonyHewins/scripts/internal/app"
)

var defaultMasterBranch = "master"

var helpText = fmt.Sprintf(`usage: gitm

Pulls from the default master branch into master and pulls it into
the current branch

Default master branch: %s`, defaultMasterBranch)

func main() {
	dir := "."
	l := app.NewLogRunner(dir, os.Stdout, os.Stderr)

	branch, err := l.CurrentGitBranch(dir)
	if err != nil {
		l.Fatal("Failed fetching git branch name: %s", err.Error())
	}

	switch len(os.Args) {
	case 0:
		l.Fatal("no args supplied")
	case 1:
	default:
		if arg := os.Args[1]; arg == "help" || arg == "--help" || arg == "-h" {
			fmt.Println(helpText)
			os.Exit(0)
		}

		l.Fatal("invalid arg: %s", os.Args[1])
	}

	if branch == defaultMasterBranch {
		l.Run("git", "pull", "origin", defaultMasterBranch)
		os.Exit(0)
	}

	l.Run("git", "checkout", defaultMasterBranch)
	l.Run("git", "pull", "origin", defaultMasterBranch)
	l.Run("git", "checkout", branch)
	l.Run("git", "merge", defaultMasterBranch)
}
