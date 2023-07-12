package cmd

import (
    "context"

    "github.com/emilybache/gildedrose-refactoring-kata/lib"
    "github.com/spf13/cobra"
    "go.uber.org/fx"
    "go.uber.org/fx/fxevent"
)

// Define all available commands for the CLI application
var cmds = map[string]lib.Command{
    "app:serve": NewServeCommand(),
    "test": NewTextTestFixtureCommand(),
}

// Return wrapped subcommands
func GetSubCommands(opt fx.Option) []*cobra.Command {
    var subCommands []*cobra.Command
    for name, command := range cmds {
        subCommands = append(subCommands, WrapSubCommand(name, command, opt))
    }
    return subCommands
}

// Wrap each subcommand and call its Setup method to return its information for Cobra library
func WrapSubCommand(name string, command lib.Command, opt fx.Option) *cobra.Command {
    wrappedCmd := &cobra.Command{
        Use:   name,
        Short: command.Short(),
        Run: func(c *cobra.Command, args []string) {
            logger := lib.GetLogger()
            opts := fx.Options(
                fx.WithLogger(func() fxevent.Logger {
                    return logger.GetFxLogger()
                }),
                fx.Invoke(command.Run()),
            )
            ctx := context.Background()
            app := fx.New(opt, opts)
            err := app.Start(ctx)
            defer app.Stop(ctx)
            if err != nil {
                logger.Fatal(err)
            }
        },
    }
    command.Setup(wrappedCmd)
    return wrappedCmd
}
