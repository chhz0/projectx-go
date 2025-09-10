package app

import (
	"context"
	"errors"

	"github.com/spf13/cobra"
)

type InitCmdFunc func(cmd *cobra.Command)
type CmdFunc func(cmd *cobra.Command, args []string)

type Command struct {
	Use   string
	Short string
	Long  string

	Version string // TODO: add version

	// Init setup the command flags and config
	Init    InitCmdFunc
	PreRun  CmdFunc
	Run     CmdFunc
	PostRun CmdFunc

	Args  cobra.PositionalArgs
	Cobra *cobra.Command

	SubCommands []*Command
}

func (c *Command) Exec(ctx context.Context) error {
	if err := c.compile(); err != nil {
		return errors.New("Command compile error")
	}

	return c.Cobra.ExecuteContext(ctx)
}

func (c *Command) compile() error {
	c.Cobra = &cobra.Command{
		Use:   c.Use,
		Short: c.Short,
		Long:  c.Long,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if c.PreRun != nil {
				c.PreRun(cmd, args)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			if c.Run != nil {
				c.Run(cmd, args)
			}
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			if c.PostRun != nil {
				c.PostRun(cmd, args)
			}
		},
		SilenceErrors:              true,
		SilenceUsage:               true,
		SuggestionsMinimumDistance: 2,
	}

	// Init flags and config
	if c.Init != nil {
		c.Init(c.Cobra)
	}

	for _, sub := range c.SubCommands {
		_ = sub.compile()
		c.Cobra.AddCommand(sub.Cobra)
	}

	return nil
}

type CommandOptions func(*Command)

func NewCommand(name string, opts ...CommandOptions) *Command {
	cmd := &Command{Use: name}

	for _, of := range opts {
		of(cmd)
	}

	return cmd
}
