package services

import (
    "github.com/emilybache/gildedrose-refactoring-kata/domains"
    "github.com/emilybache/gildedrose-refactoring-kata/lib"
    "github.com/emilybache/gildedrose-refactoring-kata/models"
)

type ItemUpdateServiceProvider struct {
    logger                          lib.Logger
    normalItemUpdateService         NormalItemUpdateService
    agedBrieItemUpdateService       AgedBrieItemUpdateService
    backstagePassItemUpdateService  BackstagePassItemUpdateService
    sulfurasItemUpdateService       SulfurasItemUpdateService
    conjuredItemUpdateService       ConjuredItemUpdateService
}

func NewItemUpdateServiceProvider(
    logger                          lib.Logger,
    normalItemUpdateService         NormalItemUpdateService,
    agedBrieItemUpdateService       AgedBrieItemUpdateService,
    backstagePassItemUpdateService  BackstagePassItemUpdateService,
    sulfurasItemUpdateService       SulfurasItemUpdateService,
    conjuredItemUpdateService       ConjuredItemUpdateService,
) domains.ItemUpdateServiceProvider {

    return ItemUpdateServiceProvider{
        logger:                          logger,
        normalItemUpdateService:         normalItemUpdateService,
        agedBrieItemUpdateService:       agedBrieItemUpdateService,
        backstagePassItemUpdateService:  backstagePassItemUpdateService,
        sulfurasItemUpdateService:       sulfurasItemUpdateService,
        conjuredItemUpdateService:       conjuredItemUpdateService,
    }
}

func (this ItemUpdateServiceProvider) GetUpdateService(item *models.Item) domains.ItemUpdateService {
    switch item.Model.Name {
        case "Aged Brie":
            return this.agedBrieItemUpdateService
        case "Backstage passes to a TAFKAL80ETC concert":
            return this.backstagePassItemUpdateService
        case "Sulfuras, Hand of Ragnaros":
            return this.sulfurasItemUpdateService
        case "Conjured Mana Cake":
            return this.conjuredItemUpdateService
        default:
            return this.normalItemUpdateService
    }
}
