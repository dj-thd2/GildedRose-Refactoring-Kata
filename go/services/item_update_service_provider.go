package services

import (
    "github.com/emilybache/gildedrose-refactoring-kata/domains"
    "github.com/emilybache/gildedrose-refactoring-kata/lib"
    "github.com/emilybache/gildedrose-refactoring-kata/models"
)

// Properties
type ItemUpdateServiceProvider struct {
    logger                          lib.Logger
    normalItemUpdateService         NormalItemUpdateService
    agedBrieItemUpdateService       AgedBrieItemUpdateService
    backstagePassItemUpdateService  BackstagePassItemUpdateService
    sulfurasItemUpdateService       SulfurasItemUpdateService
}

// Instantiate a new ItemUpdateServiceProvider
func NewItemUpdateServiceProvider(
    logger                          lib.Logger,
    normalItemUpdateService         NormalItemUpdateService,
    agedBrieItemUpdateService       AgedBrieItemUpdateService,
    backstagePassItemUpdateService  BackstagePassItemUpdateService,
    sulfurasItemUpdateService       SulfurasItemUpdateService,
) domains.ItemUpdateServiceProvider {

    return ItemUpdateServiceProvider{
        logger:                          logger,
        normalItemUpdateService:         normalItemUpdateService,
        agedBrieItemUpdateService:       agedBrieItemUpdateService,
        backstagePassItemUpdateService:  backstagePassItemUpdateService,
        sulfurasItemUpdateService:       sulfurasItemUpdateService,
    }
}

// Implement GetUpdateService method to return a different ItemUpdateService implementation depending on the item kind
func (this ItemUpdateServiceProvider) GetUpdateService(item *models.Item) domains.ItemUpdateService {
    item.Mutex.Lock()
    defer item.Mutex.Unlock()

    switch item.Model.Name {
        case "Aged Brie":
            return &this.agedBrieItemUpdateService
        case "Backstage passes to a TAFKAL80ETC concert":
            return &this.backstagePassItemUpdateService
        case "Sulfuras, Hand of Ragnaros":
            return &this.sulfurasItemUpdateService
        default:
            return &this.normalItemUpdateService
    }
}
