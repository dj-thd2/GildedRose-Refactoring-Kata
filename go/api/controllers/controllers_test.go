package controllers

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/require"

    "go.uber.org/fx"
    "go.uber.org/fx/fxtest"

    "github.com/emilybache/gildedrose-refactoring-kata/lib"
    "github.com/emilybache/gildedrose-refactoring-kata/services"
    "github.com/emilybache/gildedrose-refactoring-kata/api/middlewares"
)

// Common modules needed by the test cases instantiation for fx dependency injection
var TestCommonModules = fx.Options(
    Module,
    lib.Module,
    services.Module,
    middlewares.Module,
)

// Test case runner generic interface
type TestCaseRunner interface{}

// Set up the test case
func setupTestCase(handler lib.RequestHandler, logger lib.Logger, middleware middlewares.Middlewares, itemController ItemController, swaggerController SwaggerController) {
    logger.Info("Setup test case")
    middleware.Setup()
    swaggerController.Setup(handler.Gin)
    itemController.Setup(handler.Gin)
}

// Run a test case wrapping it with fx for dependency injection
func runTestCase(t *testing.T, runner TestCaseRunner) {
    app := fxtest.New(t, TestCommonModules, fx.Invoke(setupTestCase), fx.Invoke(runner))

    defer app.RequireStart().RequireStop()
    require.NoError(t, app.Err())
}

// Helper function to execute a HTTP request
func executeRequest(handler lib.RequestHandler, req *http.Request) *httptest.ResponseRecorder {
    result := httptest.NewRecorder()
    handler.Gin.ServeHTTP(result, req)
    return result
}
