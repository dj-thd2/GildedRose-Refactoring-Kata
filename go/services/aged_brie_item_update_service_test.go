package services

import (
    "testing"
    "github.com/stretchr/testify/assert"

    "github.com/emilybache/gildedrose-refactoring-kata/models"
)

// Aged Brie quality increments in 1 after each day
func TestAgedBrieItemUpdateService_QualityBeforeSellIn(t *testing.T) {
    runTestCase(t, func(
        agedBrieItemUpdateService AgedBrieItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Aged Brie", 5, 5})
        agedBrieItemUpdateService.UpdateQuality(item)
        assert.Equal(t, 6, item.Model.Quality)
    })
}

// If the sellIn date has been passed, the Aged Brie quality increments in 2
func TestAgedBrieItemUpdateService_QualityAfterSellIn0Days(t *testing.T) {
    runTestCase(t, func(
        agedBrieItemUpdateService AgedBrieItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Aged Brie", 0, 5})
        agedBrieItemUpdateService.UpdateQuality(item)
        assert.Equal(t, 7, item.Model.Quality)
    })
}

// If the sellIn date has been passed, the Aged Brie quality increments in 2
func TestAgedBrieItemUpdateService_QualityAfterSellIn4Days(t *testing.T) {
    runTestCase(t, func(
        agedBrieItemUpdateService AgedBrieItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Aged Brie", -4, 5})
        agedBrieItemUpdateService.UpdateQuality(item)
        assert.Equal(t, 7, item.Model.Quality)
    })
}

// Quality must not be greater than 50
func TestAgedBrieItemUpdateService_QualityNotHigherThan50(t *testing.T) {
    runTestCase(t, func(
        agedBrieItemUpdateService AgedBrieItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Aged Brie", -4, 50})
        agedBrieItemUpdateService.UpdateQuality(item)
        assert.Equal(t, 50, item.Model.Quality)
    })
}

// sellIn date must decrease
func TestAgedBrieItemUpdateService_SellInIsDecreased(t *testing.T) {
    runTestCase(t, func(
        agedBrieItemUpdateService AgedBrieItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Aged Brie", 5, 5})
        agedBrieItemUpdateService.UpdateQuality(item)
        assert.Equal(t, 4, item.Model.SellIn)
    })
}


// Test new UpdateQualityForDays method

// Item quality increments in 1 for each passed day before the sellIn date
func TestAgedBrieItemUpdateService_2Days_QualityBeforeSellIn10Days(t *testing.T) {
    runTestCase(t, func(
        agedBrieItemUpdateService AgedBrieItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Aged Brie", 10, 5})
        days := 2
        agedBrieItemUpdateService.UpdateQualityForDays(item, days)
        assert.Equal(t, 7, item.Model.Quality, "Quality should increase in 2")
        assert.Equal(t, 8, item.Model.SellIn, "SellIn should decrease in 2")
    })
}

// Item quality decrements in 1 for each passed day before the sellIn date
func TestAgedBrieItemUpdateService_10Days_QualityBeforeSellIn10Days(t *testing.T) {
    runTestCase(t, func(
        agedBrieItemUpdateService AgedBrieItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Aged Brie", 10, 20})
        days := 10
        agedBrieItemUpdateService.UpdateQualityForDays(item, days)
        assert.Equal(t, 30, item.Model.Quality, "Quality should increase in 10")
        assert.Equal(t, 0, item.Model.SellIn, "SellIn date should decrease in 10")
    })
}

// If the sellIn date has been passed, the Brie item quality increments in 2 for each extra passed day
func TestAgedBrieItemUpdateService_11Days_QualityAfterSellIn10Days(t *testing.T) {
    runTestCase(t, func(
        agedBrieItemUpdateService AgedBrieItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Aged Brie", 10, 10})
        days := 11
        agedBrieItemUpdateService.UpdateQualityForDays(item, days)
        assert.Equal(t, 28, item.Model.Quality, "Quality should increase in 10 until expired and in 2 the last day resulting in +12 total")
        assert.Equal(t, -1, item.Model.SellIn, "SellIn date should result in -1")
    })
}

// If the sellIn date has been passed, the Brie item quality increments in 2
func TestAgedBrieItemUpdateService_5Days_QualityAfterSellIn4Days(t *testing.T) {
    runTestCase(t, func(
        agedBrieItemUpdateService AgedBrieItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Aged Brie", -4, 20})
        days := 5
        agedBrieItemUpdateService.UpdateQualityForDays(item, days)
        assert.Equal(t, 30, item.Model.Quality, "Quality should increase in 10")
        assert.Equal(t, -9, item.Model.SellIn, "SellIn date should decrease in 5")
    })
}

// Quality must not be greater than 50
func TestAgedBrieItemUpdateService_3Days_QualityNotLowerThan0(t *testing.T) {
    runTestCase(t, func(
        agedBrieItemUpdateService AgedBrieItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Aged Brie", -4, 45})
        days := 3
        agedBrieItemUpdateService.UpdateQualityForDays(item, days)
        assert.Equal(t, 50, item.Model.Quality, "Quality must be 50")
        assert.Equal(t, -7, item.Model.SellIn, "SellIn date should decrease in 3")
    })
}

// sellIn date must decrease
func TestAgedBrieItemUpdateService_8Days_SellInIsDecreased(t *testing.T) {
    runTestCase(t, func(
        agedBrieItemUpdateService AgedBrieItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Aged Brie", -4, 0})
        days := 8
        agedBrieItemUpdateService.UpdateQualityForDays(item, days)
        assert.Equal(t, -12, item.Model.SellIn, "SellIn date should decrease in 8")
    })
}
