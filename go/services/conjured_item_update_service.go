package services

import (
    "github.com/emilybache/gildedrose-refactoring-kata/lib"
    "github.com/emilybache/gildedrose-refactoring-kata/models"
)

type ConjuredItemUpdateService struct {
    logger     lib.Logger
    utils      lib.Utils
}

func NewConjuredItemUpdateService(logger lib.Logger, utils lib.Utils) ConjuredItemUpdateService {
    return ConjuredItemUpdateService{
        logger:     logger,
        utils:      utils,
    }
}

func (this ConjuredItemUpdateService) UpdateQuality(item *models.Item) error {
    return this.UpdateQualityForDays(item, 1)
}

func (this ConjuredItemUpdateService) UpdateQualityForDays(item *models.Item, days int) error {
    item.Mutex.Lock()
    defer item.Mutex.Unlock()

    // The quality will decrement by 2 units in normal conditions for each day
    normalDecrement := 2

    // Calculate days to pass before SellIn date
    passedDaysBeforeSellIn := this.utils.Max(this.utils.Min(days, item.Model.SellIn), 0)

    // Calculate days to pass after SellIn date
    passedDaysAfterSellIn := days - passedDaysBeforeSellIn

    // Calculate total decrement
    decrement := (normalDecrement * passedDaysBeforeSellIn) + (2 * normalDecrement * passedDaysAfterSellIn)

    // Apply decrement to quality
    if (item.Model.Quality - decrement) < 0 {
        item.Model.Quality = 0
    } else {
        item.Model.Quality -= decrement
    }

    // Decrement sellIn date
    item.Model.SellIn -= days

    // Return no errors
    return nil
}
