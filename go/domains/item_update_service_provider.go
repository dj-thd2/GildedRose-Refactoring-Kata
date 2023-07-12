package domains

import (
    "github.com/emilybache/gildedrose-refactoring-kata/models"
)

// ItemUpdateServiceProvider interface definition
// It provides a UpdateService implementation
type ItemUpdateServiceProvider interface {
    // It should return a pointer to an instance of ItemUpdateService according to the passed Item
    GetUpdateService(item *models.Item) ItemUpdateService
}
