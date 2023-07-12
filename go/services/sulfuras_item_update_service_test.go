package services

import (
    "testing"
    "github.com/stretchr/testify/assert"

    "github.com/emilybache/gildedrose-refactoring-kata/models"
)

// Sulfuras item quality is 80 before sellIn
func TestSulfurasItemUpdateService_QualityBeforeSellIn(t *testing.T) {
    runTestCase(t, func(
        sulfurasItemUpdateService SulfurasItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Sulfuras, Hand of Ragnaros", 5, 5})
        sulfurasItemUpdateService.UpdateQuality(item)
        assert.Equal(t, 80, item.Model.Quality)
    })
}

// Sulfuras item sellIn is immutable
func TestSulfurasItemUpdateService_SellInIsImmutable(t *testing.T) {
    runTestCase(t, func(
        sulfurasItemUpdateService SulfurasItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Sulfuras, Hand of Ragnaros", 5, 5})
        sulfurasItemUpdateService.UpdateQuality(item)
        assert.Equal(t, 5, item.Model.SellIn)
    })
}

// If the sellIn date has been passed, the Sulfuras item quality is still 80
func TestSulfurasItemUpdateService_QualityAfterSellIn0Days(t *testing.T) {
    runTestCase(t, func(
        sulfurasItemUpdateService SulfurasItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Sulfuras, Hand of Ragnaros", 0, 5})
        sulfurasItemUpdateService.UpdateQuality(item)
        assert.Equal(t, 80, item.Model.Quality)
    })
}

// If the sellIn date has been passed, the Sulfuras item quality is still 80
func TestSulfurasItemUpdateService_QualityAfterSellIn4Days(t *testing.T) {
    runTestCase(t, func(
        sulfurasItemUpdateService SulfurasItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Sulfuras, Hand of Ragnaros", -4, 5})
        sulfurasItemUpdateService.UpdateQuality(item)
        assert.Equal(t, 80, item.Model.Quality)
    })
}

// Sulfuras item negative sellIn is immutable
func TestSulfurasItemUpdateService_NegativeSellInIsImmutable(t *testing.T) {
    runTestCase(t, func(
        sulfurasItemUpdateService SulfurasItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Sulfuras, Hand of Ragnaros", -4, 5})
        sulfurasItemUpdateService.UpdateQuality(item)
        assert.Equal(t, -4, item.Model.SellIn)
    })
}



// Test new UpdateQualityForDays method

// Item quality should be always 80 and sellIn immutable
func TestSulfurasItemUpdateService_2Days_QualityBeforeSellIn10Days(t *testing.T) {
    runTestCase(t, func(
        sulfurasItemUpdateService SulfurasItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Sulfuras, Hand of Ragnaros", 10, 5})
        days := 2
        sulfurasItemUpdateService.UpdateQualityForDays(item, days)
        assert.Equal(t, 80, item.Model.Quality, "Quality should be 80")
        assert.Equal(t, 10, item.Model.SellIn, "SellIn should be immutable")
    })
}

// Item quality should be always 80 and sellIn immutable
func TestSulfurasItemUpdateService_10Days_QualityBeforeSellIn10Days(t *testing.T) {
    runTestCase(t, func(
        sulfurasItemUpdateService SulfurasItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Sulfuras, Hand of Ragnaros", 10, 40})
        days := 10
        sulfurasItemUpdateService.UpdateQualityForDays(item, days)
        assert.Equal(t, 80, item.Model.Quality, "Quality should be 80")
        assert.Equal(t, 10, item.Model.SellIn, "SellIn should be immutable")
    })
}

// Item quality should be always 80 and sellIn immutable
func TestSulfurasItemUpdateService_11Days_QualityAfterSellIn10Days(t *testing.T) {
    runTestCase(t, func(
        sulfurasItemUpdateService SulfurasItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Sulfuras, Hand of Ragnaros", 10, 40})
        days := 11
        sulfurasItemUpdateService.UpdateQualityForDays(item, days)
        assert.Equal(t, 80, item.Model.Quality, "Quality should be 80")
        assert.Equal(t, 10, item.Model.SellIn, "SellIn should be immutable")
    })
}

// Item quality should be always 80 and sellIn immutable
func TestSulfurasItemUpdateService_5Days_QualityAfterSellIn4Days(t *testing.T) {
    runTestCase(t, func(
        sulfurasItemUpdateService SulfurasItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Sulfuras, Hand of Ragnaros", -4, 40})
        days := 5
        sulfurasItemUpdateService.UpdateQualityForDays(item, days)
        assert.Equal(t, 80, item.Model.Quality, "Quality should be 80")
        assert.Equal(t, -4, item.Model.SellIn, "SellIn should be immutable")
    })
}
