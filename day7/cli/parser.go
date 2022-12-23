package cli

import (
	"strings"

	"github.com/thesepehrm/aoc2022/day7/fs"
)

type parserState uint8

const (
	idle parserState = iota
	ls
)

const inputSymbol = "$"

type CommandLine struct {
	state parserState
	Fs    *fs.FileSystem
}

func NewCommandLine() *CommandLine {
	return &CommandLine{
		state: idle,
		Fs:    fs.NewFileSystem(),
	}
}

func (cli *CommandLine) Parse(c string) {
	if string(c[0]) == inputSymbol {
		a := strings.Split(c, " ")
		cli.parseArgs(a[1:])
		return
	}

	cli.parseOutput(c)
}

func (cli *CommandLine) parseArgs(a []string) {
	command := a[0]
	switch command {
	case "ls":
		cli.state = ls
	case "cd":
		cli.Fs.ChangeDir(a[1])
	}

}

func (cli *CommandLine) parseOutput(o string) {
	if string(o[0]) == "d" {
		fs.NewDirFromOutput(o, cli.Fs.CurrentDir)
		return
	}

	fs.NewFileFromOutput(o, cli.Fs.CurrentDir)
}
