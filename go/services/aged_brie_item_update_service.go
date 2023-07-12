package services

import (
    "github.com/emilybache/gildedrose-refactoring-kata/lib"
    "github.com/emilybache/gildedrose-refactoring-kata/models"
)

type AgedBrieItemUpdateService struct {
    logger     lib.Logger
    utils      lib.Utils
}

func NewAgedBrieItemUpdateService(logger lib.Logger, utils lib.Utils) AgedBrieItemUpdateService {
    return AgedBrieItemUpdateService{
        logger:   logger,
        utils:    utils,
    }
}

func (this AgedBrieItemUpdateService) UpdateQuality(item *models.Item) error {
    return this.UpdateQualityForDays(item, 1)
}

func (this AgedBrieItemUpdateService) UpdateQualityForDays(item *models.Item, days int) error {
    item.Mutex.Lock()
    defer item.Mutex.Unlock()

    // The quality will increment by 1 unit in normal conditions for each day
    normalIncrement := 1

    // Calculate days to pass before SellIn date
    passedDaysBeforeSellIn := this.utils.Max(this.utils.Min(days, item.Model.SellIn), 0)

    // Calculate days to pass after SellIn date
    passedDaysAfterSellIn := days - passedDaysBeforeSellIn

    // Calculate total decrement
    increment := (normalIncrement * passedDaysBeforeSellIn) + (2 * normalIncrement * passedDaysAfterSellIn)

    // Apply increment to quality
    if (item.Model.Quality + increment) > 50 {
        item.Model.Quality = 50
    } else {
        item.Model.Quality += increment
    }

    // Decrement sellIn date
    item.Model.SellIn -= days

    // Return no errors
    return nil
}
