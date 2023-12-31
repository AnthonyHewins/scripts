package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/AnthonyHewins/scripts/internal/app"
)

var helpText = `usage: gitp [COMMAND | <commit message>]
where COMMAND is:

h, help				Print this help
-bn					Use the git branch for the commit message
`

func main() {
	dir := "."

	l := app.NewLogRunner(dir, os.Stdout, os.Stderr)

	switch len(os.Args) {
	case 0:
		l.Fatal("must supply arguments for commit message")
	case 1:
		// no op
	default:
		push(l, os.Args[1:]...)
	}

	commitMsg := os.Args[0]
	switch commitMsg {
	case "h", "help":
		fmt.Println(helpText)
		os.Exit(0)
	case "-bn":
		branch, err := app.CurrentGitBranch(dir)
		if err != nil {
			l.Fatal("Failed fetching git branch name: %s", err.Error())
		}

		commitMsg = branch
	}

	push(l, commitMsg)
}

func push(l *app.LogRunner, msg ...string) {
	l.Run("git", "status")
	l.Run("git", "add", "-A")
	l.Run("git", "commit", "-m", strings.Join(msg, " "))
	l.Run("git", "push", "origin", "master")
	os.Exit(0)
}
