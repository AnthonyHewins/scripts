package app

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

var (
	infoColor     = color.New(color.FgCyan)
	infoBoldColor = color.New(color.FgWhite, color.BgCyan, color.Bold)

	errColor = color.New(color.FgWhite, color.Bold, color.BgRed)
	//	dirColor      = color.New(color.FgWhite, color.Bold, color.BgMagenta)
	//	templateColor = color.New(color.FgWhite, color.Bold, color.BgHiBlue)
	commandTitleColor = color.New(color.FgWhite, color.Bold, color.BgGreen)
	commandTextColor  = color.New(color.FgGreen)
)

type LogRunner struct {
	dir                      string
	logExporter, errExporter io.Writer
}

func NewLogRunner(dir string, logExporter io.Writer, errExporter io.Writer) *LogRunner {
	return &LogRunner{
		logExporter: logExporter,
		errExporter: errExporter,
	}
}

// denotes a group of steps in the server creation process
func (s *LogRunner) Info(str string, args ...any) {
	fmt.Fprintf(
		s.logExporter,
		infoColor.Sprintf(str, args...)+"\n",
	)
}

// denotes a group of steps in the server creation process
func (s *LogRunner) InfoBold(str string, args ...any) {
	fmt.Fprintf(
		s.logExporter,
		infoBoldColor.Sprintf(str, args...)+"\n",
	)
}

func (s *LogRunner) Fatal(str string, args ...any) {
	fmt.Fprintf(
		s.errExporter,
		errColor.Sprintf(str, args...)+"\n",
	)

	os.Exit(1)
}

func (s *LogRunner) Run(cmd string, args ...string) {
	fmt.Fprintf(
		s.logExporter,
		commandTitleColor.Sprint("  CMD   ")+commandTextColor.Sprintf(" %s %s\n", cmd, strings.Join(args, " ")),
	)

	command := exec.Command(cmd, args...)
	command.Dir = s.dir
	buf, err := command.Output()
	if err != nil {
		s.Fatal("failed running command: %v. Your app should still be built, but gofast couldn't do all the work for you", err)
	}

	fmt.Fprint(s.logExporter, string(buf))
}
