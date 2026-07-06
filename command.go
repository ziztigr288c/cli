package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Command is a subcommand for a cli.App.
type Command struct {
	// The name of the command
	Name string
	// A list of aliases for the command
	Aliases []string
	// A short description of the usage of this command
	Usage string
	// Custom text to show on USAGE section of help
	UsageText string
	// A longer explanation of how the command works
	Description string
	// The category the command exists in on the help page
	Category string
	// The function to call when this command is invoked
	Action ActionFunc
	// An action to execute before this command is run.
	Before BeforeFunc
	// An action to execute after this command is run.
	After AfterFunc
	// The list of flags to parse for this command
	Flags []Flag
	// The list of subcommands to parse for this command
	Subcommands []*Command
	// Whether to skip flag parsing for this command
	SkipFlagParsing bool
	// Whether to hide this command in the help text
	HideHelp bool
	// Whether to hide this command in the help text aliases
	HideHelpCommand bool
	// The name of the helper command
	HelpName string
	// CustomHelpTemplate the text template for the command help of this command
	CustomHelpTemplate string

	// OnUsageError is an action to execute when a usage error occurs
	OnUsageError OnUsageErrorFunc

	// ArgsUsage provides a description of the arguments for this command
	ArgsUsage string
}

// Run runs the command with the given context.
func (c *Command) Run(cCtx *Context) (err error) {
	if len(c.Subcommands) > 0 {
		return c.startSubcommand(cCtx)
	}

	if c.Before != nil {
		if err := c.Before(cCtx); err != nil {
			return err
		}
	}

	if c.Action != nil {
		if err := c.Action(cCtx); err != nil {
			return err
		}
	}

	if c.After != nil {
		if err := c.After(cCtx); err != nil {
			return err
		}
	}

	return nil
}

func (c *Command) startSubcommand(cCtx *Context) error {
	args := cCtx.Args()
	if args.Present() {
		name := args.First()
		// Check if '--' was encountered before this argument in the original arguments.
		// If so, we must treat it strictly as a positional argument and not run it as a subcommand.
		doubleDashIdx := -1
		for i, arg := range cCtx.App.args {
			if arg == "--" {
				doubleDashIdx = i
				break
			}
		}
		
		// We need to map the current argument back to the original args to see if it's after '--'
		isAfterDoubleDash := false
		if doubleDashIdx != -1 {
			// Find the index of the current argument in the original args
			// Since we parse sequentially, we can track if we have passed '--'
			// A simpler way: if the parent context or current context args slice has '--' in its history.
			// Let's check if '--' is in the arguments that were parsed to get here.
		}

		// Actually, a cleaner way is to check if '--' is present in the arguments passed to the current command.
		// When we parse flags, if '--' is present, we split the arguments.
		// Let's look at how parseFlags is implemented.
	}
	return nil
}