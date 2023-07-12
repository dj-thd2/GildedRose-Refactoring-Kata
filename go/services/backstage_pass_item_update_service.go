package services

import (
    "github.com/emilybache/gildedrose-refactoring-kata/lib"
    "github.com/emilybache/gildedrose-refactoring-kata/models"
)

type BackstagePassItemUpdateService struct {
    logger     lib.Logger
}

func NewBackstagePassItemUpdateService(logger lib.Logger) BackstagePassItemUpdateService {
    return BackstagePassItemUpdateService{
        logger:     logger,
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
    return nil
}
