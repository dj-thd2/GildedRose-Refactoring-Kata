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

// Test /update_quality with valid JSON in POST data should return 200
func TestUpdateQuality(t *testing.T) {
    runTestCase(t, func(
        handler lib.RequestHandler,
    ) {
        req, _ := http.NewRequest("POST", "/update_quality", strings.NewReader(`
        { "days": 30, "items": [
            { "name": "+5 Dexterity Vest", "sellIn": 10, "quality": 20 },
            { "name": "Aged Brie", "sellIn": 2, "quality": 0 },
            { "name": "Elixir of the Mongoose", "sellIn": 5, "quality": 7 },
            { "name": "Sulfuras, Hand of Ragnaros", "sellIn": 0, "quality": 80 },
            { "name": "Sulfuras, Hand of Ragnaros", "sellIn": -1, "quality": 80 },
            { "name": "Backstage passes to a TAFKAL80ETC concert", "sellIn": 15, "quality": 20 },
            { "name": "Backstage passes to a TAFKAL80ETC concert", "sellIn": 10, "quality": 49 },
            { "name": "Backstage passes to a TAFKAL80ETC concert", "sellIn": 5, "quality": 49 },
            { "name": "Conjured Mana Cake", "sellIn": 3, "quality": 6 }
        ]}
        `))
        req.Header = map[string][]string{"Content-Type": {"application/json"}}

        w := executeRequest(handler, req)

        expectedResultJson := `[
            { "name": "+5 Dexterity Vest", "sellIn": -20, "quality": 0 },
            { "name": "Aged Brie", "sellIn": -28, "quality": 50 },
            { "name": "Elixir of the Mongoose", "sellIn": -25, "quality": 0 },
            { "name": "Sulfuras, Hand of Ragnaros", "sellIn": 0, "quality": 80 },
            { "name": "Sulfuras, Hand of Ragnaros", "sellIn": -1, "quality": 80 },
            { "name": "Backstage passes to a TAFKAL80ETC concert", "sellIn": -15, "quality": 0 },
            { "name": "Backstage passes to a TAFKAL80ETC concert", "sellIn": -20, "quality": 0 },
            { "name": "Backstage passes to a TAFKAL80ETC concert", "sellIn": -25, "quality": 0 },
            { "name": "Conjured Mana Cake", "sellIn": -27, "quality": 0 }
        ]`

        assert.Equal(t, 200, w.Code)
        assert.JSONEq(t, expectedResultJson, string(w.Body.Bytes()))
    })
}
