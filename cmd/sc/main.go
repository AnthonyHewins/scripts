package main

import (
	"os"
	"strings"

	"github.com/AnthonyHewins/scripts/internal/app"
)

const helpText = `usage: sc [BRANCH TYPE] NAME...

BRANCH TYPE						f - feature
								t - task
								b - bug fix
								h - hot fix

NAME							Branch name that will be formatted`

func main() {
	l := app.NewLogRunner(".", os.Stdout, os.Stderr)

	args := l.ParseArgs(&app.Arg{HelpText: helpText})
	n := len(args)
	if n < 2 {
		l.Fatal("not enough args")
	}

	branch := ""
	switch args[0] {
	case "f":
		branch = "feature"
	case "b":
		branch = "bugfix"
	case "h":
		branch = "hotfix"
	case "t":
		branch = "task"
	default:
		l.Fatal("invalid branch type: %s", args[0])
	}

	l.Run("git", "checkout", branch+"/"+strings.Join(args[1:], "-"))
}
