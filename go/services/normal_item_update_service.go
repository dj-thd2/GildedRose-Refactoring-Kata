package services

import (
    "github.com/emilybache/gildedrose-refactoring-kata/lib"
    "github.com/emilybache/gildedrose-refactoring-kata/models"
)

type NormalItemUpdateService struct {
    logger     lib.Logger
}

func NewNormalItemUpdateService(logger lib.Logger) NormalItemUpdateService {
    return NormalItemUpdateService{
        logger:     logger,
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
