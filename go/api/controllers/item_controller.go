package controllers

import (
    "net/http"

    "github.com/emilybache/gildedrose-refactoring-kata/domains"
    "github.com/emilybache/gildedrose-refactoring-kata/models"
    "github.com/emilybache/gildedrose-refactoring-kata/lib"
    "github.com/emilybache/gildedrose-refactoring-kata/api/middlewares"
    "github.com/gin-gonic/gin"
)

// ItemController properties
type ItemController struct {
    logger                       lib.Logger
    contentTypeFilterMiddleware  middlewares.ContentTypeFilterMiddleware
    itemUpdateServiceProvider    domains.ItemUpdateServiceProvider
}

// Model for POST request for /update_quality endpoint
type UpdateQualityRequest struct {
    Days   int                    `json:"days" binding:"required,gte=1"`
    Items  []*models.ItemModel    `json:"items" binding:"required"`
}

// Instantiate an ItemController
func NewItemController(
    logger                       lib.Logger,
    contentTypeFilterMiddleware  middlewares.ContentTypeFilterMiddleware,
    itemUpdateServiceProvider    domains.ItemUpdateServiceProvider,
) ItemController {
    return ItemController{
        logger:                       logger,
        contentTypeFilterMiddleware:  contentTypeFilterMiddleware,
        itemUpdateServiceProvider:    itemUpdateServiceProvider,
    }
}

// Setup ItemController routes
func (this *ItemController) Setup(engine *gin.Engine) {
    engine.GET("/status", this.getStatus)
    engine.POST("/update_quality", this.contentTypeFilterMiddleware.Handler("application/json"), this.postUpdateQuality)
    engine.NoRoute(this.invalidPath)
}


// This method will be called for all the paths not covered by any route to return a 404 Not Found
func (this *ItemController) invalidPath(ctx *gin.Context) {
    ctx.AbortWithStatus(http.StatusNotFound)
}

// API Endpoints

// getStatus godoc
// @Summary Health check (get status)
// @Schemes
// @Description Return a simple OK response
// @Tags health
// @Accept json
// @Produce json
// @Success 200 "Service is up and running"
// @Router /status [get]
func (this *ItemController) getStatus(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{
        "status": "ok",
    })
}

// postUpdateQuality godoc
// @Summary Execute Update Quality for all given items and passed days
// @Schemes
// @Description Load a full list of items that will be handled by the service and process them, returning the item list with the updated values for the given days
// @Tags items
// @Accept json
// @Param message body UpdateQualityRequest true "Items info"
// @Success 200 "Result of items with updated quality"
// @Failure 400 "Invalid data in the request"
// @Router /update_quality [post]
func (this *ItemController) postUpdateQuality(ctx *gin.Context) {

    // Request data
    var request UpdateQualityRequest

    if err := ctx.BindJSON(&request); err != nil {
        ctx.AbortWithStatus(http.StatusBadRequest)
        return
    }

    // Instantiate Item objects from the items got from the request
    // Item objects have an ItemModel object with the actual data and a Mutex to synchronize concurrent access
    // As we use pointers, it should not introduce too much overhead
    var items = make([]*models.Item, len(request.Items))
    for i, item := range request.Items {
        items[i] = models.NewItem(item)
    }

    // Call UpdateQuality for each item of the array, as many times as defined in request.Days
    // For performance reasons we implemented a method to directly calculate the final values instead of cycle through each day
    for _, item := range items {
        // Get ItemUpdateService for that specific item and call its UpdateQuality method with the specified number of days
        itemUpdateService := this.itemUpdateServiceProvider.GetUpdateService(item)
        itemUpdateService.UpdateQualityForDays(item, request.Days)
    }

    // Generate a ItemModel array of the objects with the actual data, to be returned as a reply
    // As we use pointers, it should not introduce too much overhead
    var itemEntities = make([]*models.ItemModel, len(items))
    for i, item := range items {
        itemEntities[i] = item.Model
    }

    ctx.JSON(http.StatusOK, itemEntities)
}
