package cmd

import (
	"errors"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var errCommandNotFound = errors.New("command not found")

// App represents a application
type App struct {
	*kingpin.Application
	cmds map[string]Runner
}

// Runner implements a function to run commands
type Runner interface {
	Run() error
}

// NewApp creates a new application
func NewApp(name, desc string) *App {
	return &App{
		Application: kingpin.New(name, desc),
		cmds:        make(map[string]Runner),
	}
}

// RegisterCommand registers a runner for a given full command
func (a *App) RegisterCommand(fullCmd string, r Runner) {
	a.cmds[fullCmd] = r
}

// Run executes a command
func (a *App) Run() error {
	cmd, err := a.Application.Parse(os.Args[1:])
	if err != nil {
		return err
	}
	if a.cmds[cmd] == nil {
		return errCommandNotFound
	}
	return a.cmds[cmd].Run()
}
