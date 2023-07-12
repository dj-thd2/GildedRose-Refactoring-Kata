package models

import (
    "fmt"
    "sync"
)

// Plain Item object data to be marshalled/unmarshalled from JSON
type ItemModel struct {
    Name     string `json:"name" binding:"required"`
    SellIn   int    `json:"sellIn" binding:"required"`
    Quality  int    `json:"quality" binding:"required,gte=0"`
}

// Loaded Item object with the ItemModel and a Mutex to synchronize concurrent access
type Item struct {
    Model   *ItemModel
    Mutex   *sync.Mutex
}

// Create a new Item pointing to a ItemModel instance, creating its Mutex object
func NewItem(model *ItemModel) *Item {
    return &Item{
        Model: model,
        Mutex: &sync.Mutex{},
    }
}

// Return String from Item object
func (this Item) String() string {
    this.Mutex.Lock()
    defer this.Mutex.Unlock()

    return this.Model.String()
}

// Return String from ItemModel object
func (this ItemModel) String() string {
    return fmt.Sprintf("%s, %d, %d", this.Name, this.SellIn, this.Quality)
}
