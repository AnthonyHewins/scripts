package main

import (
	"os"
	"strings"

	"github.com/AnthonyHewins/scripts/internal/app"
)

var helpText = `usage: gitp [COMMAND | <commit message>]
where COMMAND is:

h, help				Print this help
-bn					Use the git branch for the commit message

Anything else is interpreted as a commit message
`

func main() {
	dir := "."

	l := app.NewLogRunner(dir, os.Stdout, os.Stderr)

	args := l.ParseArgs(&app.Arg{HelpText: helpText})

	commitMsg := ""
	switch args[1] {
	case "-bn":
		branch, err := l.CurrentGitBranch(dir)
		if err != nil {
			l.Fatal("Failed fetching git branch name: %s", err.Error())
		}

		commitMsg = branch
	default:
		commitMsg = "'" + strings.Join(os.Args[1:], " ") + "'"
	}

	gitBranch, err := l.CurrentGitBranch(".")
	if err != nil {
		panic(err)
	}

	l.Run("git", "status")
	l.Run("git", "add", "-A")
	l.Run("git", "commit", "-m", commitMsg)
	l.Run("git", "push", "origin", gitBranch)
}
