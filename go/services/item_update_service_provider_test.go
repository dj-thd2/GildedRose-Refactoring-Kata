package services

import (
    "testing"
    "github.com/stretchr/testify/assert"

    "github.com/emilybache/gildedrose-refactoring-kata/models"
    "github.com/emilybache/gildedrose-refactoring-kata/domains"
)

// Normal item should use a NormalItemUpdateService
func TestItemUpdateServiceProvider_NormalItem(t *testing.T) {
    runTestCase(t, func(
        itemUpdateServiceProvider domains.ItemUpdateServiceProvider,
    ) {
        item := models.NewItem(&models.ItemModel{"Random normal item", 5, 5})
        updateServiceProvider := itemUpdateServiceProvider.GetUpdateService(item);
        _, ok := updateServiceProvider.(*NormalItemUpdateService)
        assert.Equal(t, true, ok)
    })
}

// Aged brie item should use a AgedBrieItemUpdateService
func TestItemUpdateServiceProvider_AgedBrieItem(t *testing.T) {
    runTestCase(t, func(
        itemUpdateServiceProvider domains.ItemUpdateServiceProvider,
    ) {
        item := models.NewItem(&models.ItemModel{"Aged Brie", 5, 5})
        updateServiceProvider := itemUpdateServiceProvider.GetUpdateService(item);
        _, ok := updateServiceProvider.(*AgedBrieItemUpdateService)
        assert.Equal(t, true, ok)
    })
}

// Backstage pass item should use a BackstagePassItemUpdateService
func TestItemUpdateServiceProvider_BackstagePassItem(t *testing.T) {
    runTestCase(t, func(
        itemUpdateServiceProvider domains.ItemUpdateServiceProvider,
    ) {
        item := models.NewItem(&models.ItemModel{"Backstage passes to a TAFKAL80ETC concert", 5, 5})
        updateServiceProvider := itemUpdateServiceProvider.GetUpdateService(item);
        _, ok := updateServiceProvider.(*BackstagePassItemUpdateService)
        assert.Equal(t, true, ok)
    })
}

// Sulfuras item should use a SulfurasItemUpdateService
func TestItemUpdateServiceProvider_SulfurasItem(t *testing.T) {
    runTestCase(t, func(
        itemUpdateServiceProvider domains.ItemUpdateServiceProvider,
    ) {
        item := models.NewItem(&models.ItemModel{"Sulfuras, Hand of Ragnaros", 5, 5})
        updateServiceProvider := itemUpdateServiceProvider.GetUpdateService(item);
        _, ok := updateServiceProvider.(*SulfurasItemUpdateService)
        assert.Equal(t, true, ok)
    })
}

func TestItemUpdateServiceProvider_ConjuredItem(t *testing.T) {
    runTestCase(t, func(
        itemUpdateServiceProvider domains.ItemUpdateServiceProvider,
    ) {
        item := models.NewItem(&models.ItemModel{"Conjured Mana Cake", 5, 5})
        updateServiceProvider := itemUpdateServiceProvider.GetUpdateService(item);
        _, ok := updateServiceProvider.(*ConjuredItemUpdateService)
        assert.Equal(t, true, ok)
    })
}
