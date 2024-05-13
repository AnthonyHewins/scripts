package app

import "os"

type Arg struct {
	HelpText string
}

func (l *LogRunner) ParseArgs(a *Arg) []string {
	n := len(os.Args)

	switch n {
	case 0:
		l.Fatal("missing args")
	case 2:
		arg := os.Args[1]
		if arg == "help" || arg == "-h" || arg == "--help" {
			l.logExporter.Write([]byte(a.HelpText))
			os.Exit(0)
		}
	}

	return os.Args
}
