package domains

import (
    "github.com/emilybache/gildedrose-refactoring-kata/models"
)

// ItemUpdateService interface definition
// It implements the logic to update a item
type ItemUpdateService interface {
    // Update item quality for 1 day
    UpdateQuality(item *models.Item) error
}
