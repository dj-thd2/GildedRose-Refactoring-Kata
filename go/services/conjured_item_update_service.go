package services

import (
    "github.com/emilybache/gildedrose-refactoring-kata/lib"
    "github.com/emilybache/gildedrose-refactoring-kata/models"
)

type ConjuredItemUpdateService struct {
    logger     lib.Logger
}

func NewConjuredItemUpdateService(logger lib.Logger) ConjuredItemUpdateService {
    return ConjuredItemUpdateService{
        logger:     logger,
    }
}

func (this ConjuredItemUpdateService) UpdateQuality(item *models.Item) error {
    item.Mutex.Lock()
    defer item.Mutex.Unlock()

    // The quality will decrement by 2 unit in normal conditions (double rate than normal)
    decrement := 2

    // If the item sellIn date is due, the quality will decrement at double rate
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
