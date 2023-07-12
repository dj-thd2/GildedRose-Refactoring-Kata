package services

import (
    "github.com/emilybache/gildedrose-refactoring-kata/lib"
    "github.com/emilybache/gildedrose-refactoring-kata/models"
)

type NormalItemUpdateService struct {
    logger     lib.Logger
    utils      lib.Utils
}

func NewNormalItemUpdateService(logger lib.Logger, utils lib.Utils) NormalItemUpdateService {
    return NormalItemUpdateService{
        logger:     logger,
        utils:      utils,
    }
}

func (this NormalItemUpdateService) UpdateQuality(item *models.Item) error {
    item.Mutex.Lock()
    defer item.Mutex.Unlock()

    // The quality will decrement by 1 unit in normal conditions
    decrement := 1

    // If the sellIn date is due, the decrement will be double than normal
    if item.Model.SellIn <= 0 {
        decrement *= 2
    }

    // Decrement quality and force it to 0 if the result is negative
    if (item.Model.Quality - decrement) < 0 {
        item.Model.Quality = 0
    } else {
        item.Model.Quality -= decrement
    }

    // Decrement sellIn date
    item.Model.SellIn--

    // Return no errors
    return nil
}

func (this NormalItemUpdateService) UpdateQualityForDays(item *models.Item, days int) error {
    item.Mutex.Lock()
    defer item.Mutex.Unlock()

    // The quality will decrement by 1 unit in normal conditions for each day
    normalDecrement := 1

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
