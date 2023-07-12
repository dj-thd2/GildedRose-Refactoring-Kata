package controllers

import (
    "net/http"
    "testing"
    "strings"

    "github.com/stretchr/testify/assert"

    "github.com/emilybache/gildedrose-refactoring-kata/lib"
)

// Test /status should return 200
func TestStatus(t *testing.T) {
    runTestCase(t, func(
        handler lib.RequestHandler,
    ) {
        req, _ := http.NewRequest("GET", "/status", nil)

        w := executeRequest(handler, req)
        assert.Equal(t, 200, w.Code)
        assert.Equal(t, `{"status":"ok"}`, w.Body.String())
    })
}

// Test /notfoundurl should return 404
func TestNotFound(t *testing.T) {
    runTestCase(t, func(
        handler lib.RequestHandler,
    ) {
        req, _ := http.NewRequest("GET", "/notfoundurl", nil)

        w := executeRequest(handler, req)
        assert.Equal(t, 404, w.Code)
    })
}

// Test /update_quality with invalid JSON in POST data should return 400
func TestInvalidItem(t *testing.T) {
    runTestCase(t, func(
        handler lib.RequestHandler,
    ) {
        req, _ := http.NewRequest("POST", "/update_quality", strings.NewReader(`
        { "id": 1, "people": invalid
        `))
        req.Header = map[string][]string{"Content-Type": {"application/json"}}

        w := executeRequest(handler, req)
        assert.Equal(t, 400, w.Code)
    })
}

// Test /update_quality with valid JSON in POST data should return 200
// We skip unit testing for the logic as it should be already covered in service unit tests
func TestValidItem(t *testing.T) {
    runTestCase(t, func(
        handler lib.RequestHandler,
    ) {
        req, _ := http.NewRequest("POST", "/update_quality", strings.NewReader(`
        { "days": 1, "items": [{ "name": "Random normal item", "quality": 5, "sellIn": 5}]}
        `))
        req.Header = map[string][]string{"Content-Type": {"application/json"}}

        w := executeRequest(handler, req)
        assert.Equal(t, 200, w.Code)
    })
}
