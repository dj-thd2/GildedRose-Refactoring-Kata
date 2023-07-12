package services

import (
    "testing"
    "github.com/stretchr/testify/assert"

    "github.com/emilybache/gildedrose-refactoring-kata/models"
)

// Test original UpdateQuality method

// Normal item quality decrements in 1 after each day
func TestNormalItemUpdateService_QualityBeforeSellIn(t *testing.T) {
    runTestCase(t, func(
        normalItemUpdateService NormalItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Random normal item", 5, 5})
        normalItemUpdateService.UpdateQuality(item)
        assert.Equal(t, 4, item.Model.Quality)
    })
}

// If the sellIn date has been passed, the Normal item quality decrements in 2
func TestNormalItemUpdateService_QualityAfterSellIn0Days(t *testing.T) {
    runTestCase(t, func(
        normalItemUpdateService NormalItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Random normal item", 0, 5})
        normalItemUpdateService.UpdateQuality(item)
        assert.Equal(t, 3, item.Model.Quality)
    })
}

// If the sellIn date has been passed, the Normal item quality decrements in 2
func TestNormalItemUpdateService_QualityAfterSellIn4Days(t *testing.T) {
    runTestCase(t, func(
        normalItemUpdateService NormalItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Random normal item", -4, 5})
        normalItemUpdateService.UpdateQuality(item)
        assert.Equal(t, 3, item.Model.Quality)
    })
}

// Quality must not be lower than 0
func TestNormalItemUpdateService_QualityNotLowerThan0(t *testing.T) {
    runTestCase(t, func(
        normalItemUpdateService NormalItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Random normal item", -4, 0})
        normalItemUpdateService.UpdateQuality(item)
        assert.Equal(t, 0, item.Model.Quality)
    })
}

// sellIn date must decrease
func TestNormalItemUpdateService_SellInIsDecreased(t *testing.T) {
    runTestCase(t, func(
        normalItemUpdateService NormalItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Random normal item", -4, 0})
        normalItemUpdateService.UpdateQuality(item)
        assert.Equal(t, -5, item.Model.SellIn)
    })
}


// Test new UpdateQualityForDays method

// Item quality decrements in 1 for each passed day before the sellIn date
func TestNormalItemUpdateService_2Days_QualityBeforeSellIn10Days(t *testing.T) {
    runTestCase(t, func(
        normalItemUpdateService NormalItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Random normal item", 10, 5})
        days := 2
        normalItemUpdateService.UpdateQualityForDays(item, days)
        assert.Equal(t, 3, item.Model.Quality, "Quality should decrease in 2")
        assert.Equal(t, 8, item.Model.SellIn, "SellIn should decrease in 2")
    })
}

// Item quality decrements in 1 for each passed day before the sellIn date
func TestNormalItemUpdateService_10Days_QualityBeforeSellIn10Days(t *testing.T) {
    runTestCase(t, func(
        normalItemUpdateService NormalItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Random normal item", 10, 40})
        days := 10
        normalItemUpdateService.UpdateQualityForDays(item, days)
        assert.Equal(t, 30, item.Model.Quality, "Quality should decrease in 10")
        assert.Equal(t, 0, item.Model.SellIn, "SellIn date should decrease in 10")
    })
}

// If the sellIn date has been passed, the Normal item quality decrements in 2 for each extra passed day
func TestNormalItemUpdateService_11Days_QualityAfterSellIn10Days(t *testing.T) {
    runTestCase(t, func(
        normalItemUpdateService NormalItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Random normal item", 10, 40})
        days := 11
        normalItemUpdateService.UpdateQualityForDays(item, days)
        assert.Equal(t, 28, item.Model.Quality, "Quality should decrease in 10 until expired and in 2 the last day resulting in -12 total")
        assert.Equal(t, -1, item.Model.SellIn, "SellIn date should result in -1")
    })
}

// If the sellIn date has been passed, the Normal item quality decrements in 2
func TestNormalItemUpdateService_5Days_QualityAfterSellIn4Days(t *testing.T) {
    runTestCase(t, func(
        normalItemUpdateService NormalItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Random normal item", -4, 40})
        days := 5
        normalItemUpdateService.UpdateQualityForDays(item, days)
        assert.Equal(t, 30, item.Model.Quality, "Quality should decrease in 10")
        assert.Equal(t, -9, item.Model.SellIn, "SellIn date should decrease in 5")
    })
}

// Quality must not be lower than 0
func TestNormalItemUpdateService_3Days_QualityNotLowerThan0(t *testing.T) {
    runTestCase(t, func(
        normalItemUpdateService NormalItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Random normal item", -4, 5})
        days := 3
        normalItemUpdateService.UpdateQualityForDays(item, days)
        assert.Equal(t, 0, item.Model.Quality, "Quality must be 0")
        assert.Equal(t, -7, item.Model.SellIn, "SellIn date should decrease in 3")
    })
}

// sellIn date must decrease
func TestNormalItemUpdateService_8Days_SellInIsDecreased(t *testing.T) {
    runTestCase(t, func(
        normalItemUpdateService NormalItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Random normal item", -4, 0})
        days := 8
        normalItemUpdateService.UpdateQualityForDays(item, days)
        assert.Equal(t, -12, item.Model.SellIn, "SellIn date should decrease in 8")
    })
}
