package services

import (
    "github.com/emilybache/gildedrose-refactoring-kata/lib"
    "github.com/emilybache/gildedrose-refactoring-kata/models"
)

type BackstagePassItemUpdateService struct {
    logger     lib.Logger
    utils      lib.Utils
}

func NewBackstagePassItemUpdateService(logger lib.Logger, utils lib.Utils) BackstagePassItemUpdateService {
    return BackstagePassItemUpdateService{
        logger:     logger,
        utils:      utils,
    }
}

func (this BackstagePassItemUpdateService) UpdateQuality(item *models.Item) error {
    return this.UpdateQualityForDays(item, 1)
}

func (this BackstagePassItemUpdateService) UpdateQualityForDays(item *models.Item, days int) error {
    item.Mutex.Lock()
    defer item.Mutex.Unlock()

    // If final sellIn date is 0 or negative, quality is set to 0
    if (item.Model.SellIn - days) < 0 {
        item.Model.Quality = 0
    } else {
        // The quality will increment by 1 unit in normal conditions for each day
        normalIncrement := 1

        // Calculate days to pass before SellIn date
        passedDaysBeforeSellIn := this.utils.Max(this.utils.Min(days, item.Model.SellIn), 0)

        // Increment 1 for each day until sellIn date
        increment := normalIncrement * passedDaysBeforeSellIn

        // Calculate final sellIn value
        finalSellIn := item.Model.SellIn - passedDaysBeforeSellIn

        // Last 10 days increments a extra unit (total 2 units)
        if finalSellIn < 10 {
            increment += normalIncrement * this.utils.Min(10 - finalSellIn, passedDaysBeforeSellIn)
        }

        // Last 5 days increments a extra unit (total 3 units)
        if finalSellIn < 5 {
            increment += normalIncrement * this.utils.Min(5 - finalSellIn, passedDaysBeforeSellIn)
        }

        // Apply increment to quality
        if (item.Model.Quality + increment) > 50 {
            item.Model.Quality = 50
        } else {
            item.Model.Quality += increment
        }
    }

    // Decrement sellIn date
    item.Model.SellIn -= days

    // Return no errors
    return nil
}
