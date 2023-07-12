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
    item.Mutex.Lock()
    defer item.Mutex.Unlock()

    // If the item sellIn date is due, the quality will be 0
    if item.Model.SellIn <= 0 {
        item.Model.Quality = 0
    } else {
        // Default increment for quality is 1
        increment := 1

        // If the item sellIn date is less than or equal to 5 days, the quality increment will be 3
        // If the item sellIn date is between 6 and 10 days, the quality increment will be 2
        if item.Model.SellIn <= 5 {
            increment = 3
        } else if item.Model.SellIn <= 10 {
            increment = 2
        }

        // Increment quality and top it to a maximum value of 50
        if (item.Model.Quality + increment) > 50 {
            item.Model.Quality = 50
        } else {
            item.Model.Quality += increment
        }
    }

    // Decrement sellIn date
    item.Model.SellIn--

    // Return no errors
    return nil
}

func (this BackstagePassItemUpdateService) UpdateQualityForDays(item *models.Item, days int) error {
    item.Mutex.Lock()
    defer item.Mutex.Unlock()

    // If final sellIn date is 0 or negative, quality is set to 0
    if (item.Model.SellIn - days) <= 0 {
        item.Model.Quality = 0
    } else {
        // The quality will increment by 1 unit in normal conditions for each day
        normalIncrement := 1

        // Calculate days to pass before SellIn date
        passedDaysBeforeSellIn := this.utils.Max(this.utils.Min(days, item.Model.SellIn), 0)

        this.logger.Error(passedDaysBeforeSellIn)

        // Increment 1 for each day until sellIn date
        increment := normalIncrement * passedDaysBeforeSellIn

        // Calculate final sellIn value
        finalSellIn := item.Model.SellIn - passedDaysBeforeSellIn

        this.logger.Error(finalSellIn)

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
