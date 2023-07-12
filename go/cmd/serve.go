package cmd

import (
    "github.com/emilybache/gildedrose-refactoring-kata/lib"
    "github.com/emilybache/gildedrose-refactoring-kata/api/middlewares"
    "github.com/emilybache/gildedrose-refactoring-kata/api/routes"
    "github.com/spf13/cobra"
)

// Main serve command

// Generic interface
type ServeCommand struct{}

// Short description of the command
func (this *ServeCommand) Short() string {
    return "serve application"
}

// Setup method called on initialization
func (this *ServeCommand) Setup(command *cobra.Command) {}

// Command handler
func (this *ServeCommand) Run() lib.CommandRunner {
    return func(
        middleware middlewares.Middlewares,
        env lib.Env,
        requestHandler lib.RequestHandler,
        route routes.Routes,
        logger lib.Logger,
    ) {
        logger.Info("Init middleware and routes")
        middleware.Setup()
        route.Setup()

        logger.Info("Running server")
        if env.ServerPort == "" {
            _ = requestHandler.Gin.Run()
        } else {
            _ = requestHandler.Gin.Run(":" + env.ServerPort)
        }
    }
}

func NewServeCommand() *ServeCommand {
    return &ServeCommand{}
}
