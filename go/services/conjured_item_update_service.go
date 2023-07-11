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
    decrement := 1

    item.Mutex.Lock()
    defer item.Mutex.Unlock()

    if item.Model.SellIn <= 0 {
        decrement *= 2
    }
    if (item.Model.Quality - decrement) < 0 {
        item.Model.Quality = 0
    } else {
        item.Model.Quality -= decrement
    }
    item.Model.SellIn--

    return nil
}
