package services

import (
    "testing"

    "github.com/stretchr/testify/require"

    "go.uber.org/fx"
    "go.uber.org/fx/fxtest"

    "github.com/emilybache/gildedrose-refactoring-kata/lib"
)

// Common modules needed by the test cases instantiation for fx dependency injection
var TestCommonModules = fx.Options(
    Module,
    lib.Module,
)

// Test case runner generic interface
type TestCaseRunner interface{}

// Set up the test case
func setupTestCase(logger lib.Logger) {
    logger.Info("Setup Service test case")
}

// Run a test case wrapping it with fx for dependency injection
func runTestCase(t *testing.T, runner TestCaseRunner) {
    app := fxtest.New(t, TestCommonModules, fx.Invoke(setupTestCase), fx.Invoke(runner))

    defer app.RequireStart().RequireStop()
    require.NoError(t, app.Err())
}
