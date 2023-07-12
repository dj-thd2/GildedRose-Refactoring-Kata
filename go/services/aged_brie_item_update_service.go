package services

import (
    "github.com/emilybache/gildedrose-refactoring-kata/lib"
    "github.com/emilybache/gildedrose-refactoring-kata/models"
)

type AgedBrieItemUpdateService struct {
    logger     lib.Logger
}

func NewAgedBrieItemUpdateService(logger lib.Logger) AgedBrieItemUpdateService {
    return AgedBrieItemUpdateService{
        logger:   logger,
    }
}

func (this AgedBrieItemUpdateService) UpdateQuality(item *models.Item) error {
    item.Mutex.Lock()
    defer item.Mutex.Unlock()

    // The quality will increment by 1 unit in normal conditions
    increment := 1

    // The quality will increment at double rate when the item has expired
    if item.Model.SellIn <= 0 {
        increment *= 2
    }

    // Increment quality and top it to a maximum value of 50
    if (item.Model.Quality + increment) > 50 {
        item.Model.Quality = 50
    } else {
        item.Model.Quality += increment
    }

    // Decrement sellIn date
    item.Model.SellIn--

    // Return no errors
    return nil
}
