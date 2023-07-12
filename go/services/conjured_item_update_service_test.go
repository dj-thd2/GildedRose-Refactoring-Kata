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
