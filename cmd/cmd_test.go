package cmd

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestRunner struct{}

func (tr *TestRunner) Run() error {
	return errors.New("fail")
}

func TestNewApp(t *testing.T) {
	assert := assert.New(t)

	app := NewApp("test", "Thiis is a new app")

	assert.NotNil(app)
	assert.NotNil(app.Application)
	assert.NotNil(app.cmds)
}

func TestApp_RegisterCommand(t *testing.T) {
	assert := assert.New(t)

	app := NewApp("test", "Thiis is a new app")

	r := &TestRunner{}

	app.RegisterCommand("test", r)

	assert.Equal(r, app.cmds["test"], "they should be equal")
}
