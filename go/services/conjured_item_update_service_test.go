package services

import (
    "testing"
    "github.com/stretchr/testify/assert"

    "github.com/emilybache/gildedrose-refactoring-kata/models"
)

// Conjured item quality decrements in 2 after each day
func TestConjuredItemUpdateService_QualityBeforeSellIn(t *testing.T) {
    runTestCase(t, func(
        conjuredItemUpdateService ConjuredItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Conjured Mana Cake", 5, 5})
        conjuredItemUpdateService.UpdateQuality(item)
        assert.Equal(t, 3, item.Model.Quality)
    })
}

// If the sellIn date has been passed by 0 days, the Conjured item quality decrements in 4
func TestConjuredItemUpdateService_QualityAfterSellIn0Days(t *testing.T) {
    runTestCase(t, func(
        conjuredItemUpdateService ConjuredItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Conjured Mana Cake", 0, 5})
        conjuredItemUpdateService.UpdateQuality(item)
        assert.Equal(t, 1, item.Model.Quality)
    })
}

// If the sellIn date has been passed (negative), the Conjured item quality decrements in 4
func TestConjuredItemUpdateService_QualityAfterSellIn4Days(t *testing.T) {
    runTestCase(t, func(
        conjuredItemUpdateService ConjuredItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Conjured Mana Cake", -4, 5})
        conjuredItemUpdateService.UpdateQuality(item)
        assert.Equal(t, 1, item.Model.Quality)
    })
}

// Quality must not be lower than 0
func TestConjuredItemUpdateService_QualityNotLowerThan0(t *testing.T) {
    runTestCase(t, func(
        conjuredItemUpdateService ConjuredItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Conjured Mana Cake", -4, 0})
        conjuredItemUpdateService.UpdateQuality(item)
        assert.Equal(t, 0, item.Model.Quality)
    })
}

// sellIn date must decrease
func TestConjuredItemUpdateService_SellInIsDecreased(t *testing.T) {
    runTestCase(t, func(
        conjuredItemUpdateService ConjuredItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Conjured Mana Cake", -4, 0})
        conjuredItemUpdateService.UpdateQuality(item)
        assert.Equal(t, -5, item.Model.SellIn)
    })
}



// Test new UpdateQualityForDays method

// Item quality decrements in 2 for each passed day before the sellIn date
func TestConjuredItemUpdateService_2Days_QualityBeforeSellIn10Days(t *testing.T) {
    runTestCase(t, func(
        conjuredItemUpdateService ConjuredItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Conjured Mana Cake", 10, 5})
        days := 2
        conjuredItemUpdateService.UpdateQualityForDays(item, days)
        assert.Equal(t, 1, item.Model.Quality, "Quality should decrease in 4")
        assert.Equal(t, 8, item.Model.SellIn, "SellIn should decrease in 2")
    })
}

// Item quality decrements in 2 for each passed day before the sellIn date
func TestConjuredItemUpdateService_10Days_QualityBeforeSellIn10Days(t *testing.T) {
    runTestCase(t, func(
        conjuredItemUpdateService ConjuredItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Conjured Mana Cake", 10, 40})
        days := 10
        conjuredItemUpdateService.UpdateQualityForDays(item, days)
        assert.Equal(t, 20, item.Model.Quality, "Quality should decrease in 20")
        assert.Equal(t, 0, item.Model.SellIn, "SellIn date should decrease in 10")
    })
}

// If the sellIn date has been passed, the Conjured item quality decrements in 4 for each extra passed day
func TestConjuredItemUpdateService_11Days_QualityAfterSellIn10Days(t *testing.T) {
    runTestCase(t, func(
        conjuredItemUpdateService ConjuredItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Conjured Mana Cake", 10, 40})
        days := 11
        conjuredItemUpdateService.UpdateQualityForDays(item, days)
        assert.Equal(t, 16, item.Model.Quality, "Quality should decrease in 20 until expired and in 4 the last day resulting in -24 total")
        assert.Equal(t, -1, item.Model.SellIn, "SellIn date should result in -1")
    })
}

// If the sellIn date has been passed, the Conjured item quality decrements in 4
func TestConjuredItemUpdateService_5Days_QualityAfterSellIn4Days(t *testing.T) {
    runTestCase(t, func(
        conjuredItemUpdateService ConjuredItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Conjured Mana Cake", -4, 40})
        days := 5
        conjuredItemUpdateService.UpdateQualityForDays(item, days)
        assert.Equal(t, 20, item.Model.Quality, "Quality should decrease in 20")
        assert.Equal(t, -9, item.Model.SellIn, "SellIn date should decrease in 5")
    })
}

// Quality must not be lower than 0
func TestConjuredItemUpdateService_3Days_QualityNotLowerThan0(t *testing.T) {
    runTestCase(t, func(
        conjuredItemUpdateService ConjuredItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Conjured Mana Cake", -4, 5})
        days := 3
        conjuredItemUpdateService.UpdateQualityForDays(item, days)
        assert.Equal(t, 0, item.Model.Quality, "Quality must be 0")
        assert.Equal(t, -7, item.Model.SellIn, "SellIn date should decrease in 3")
    })
}

// sellIn date must decrease
func TestConjuredItemUpdateService_8Days_SellInIsDecreased(t *testing.T) {
    runTestCase(t, func(
        conjuredItemUpdateService ConjuredItemUpdateService,
    ) {
        item := models.NewItem(&models.ItemModel{"Conjured Mana Cake", -4, 0})
        days := 8
        conjuredItemUpdateService.UpdateQualityForDays(item, days)
        assert.Equal(t, -12, item.Model.SellIn, "SellIn date should decrease in 8")
    })
}
