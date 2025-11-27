package cli

import (
	"go-todos/internal/config"
	"go-todos/internal/core/ports"

	"github.com/spf13/cobra"
)

type CMD struct {
	root    *cobra.Command
	service ports.AppService
}

func NewCMD(cfg *config.Config, svc ports.AppService) *CMD {
	root := &cobra.Command{
		Use:     "go-todos",
		Short:   "A simple CLI todo application",
		Version: cfg.AppVersion,
	}

	cli := CMD{
		root:    root,
		service: svc,
	}

	cli.attachCommands()
	return &cli
}

func (c *CMD) Execute() error {
	return c.root.Execute()
}

func (c *CMD) attachCommands() {
	c.root.AddCommand(c.AddTodo())
	c.root.AddCommand(c.ListTodos())
	c.root.AddCommand(c.GetTodo())
	c.root.AddCommand(c.UpdateTodo())
	c.root.AddCommand(c.DeleteTodo())
}
