package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/AnthonyHewins/scripts/internal/app"
)

var defaultMasterBranch = "master"

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

	commitMsg := args[0]
	switch commitMsg {
	case "h", "help":
		fmt.Println(helpText)
		os.Exit(0)
	case "-bn":
		branch, err := l.CurrentGitBranch(dir)
		if err != nil {
			l.Fatal("Failed fetching git branch name: %s", err.Error())
		}

		commitMsg = branch
	}

	gitBranch, err := l.CurrentGitBranch(".")
	if err != nil {
		panic(err)
	}

	push(l, gitBranch, commitMsg)
}

func push(l *app.LogRunner, branch string, msg ...string) {
	l.Run("git", "status")
	l.Run("git", "add", "-A")
	l.Run("git", "commit", "-m", strings.Join(msg, " "))
	l.Run("git", "push", "origin", branch)
}
