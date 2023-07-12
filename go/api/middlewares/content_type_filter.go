package middlewares

import (
    "net/http"
    "github.com/emilybache/gildedrose-refactoring-kata/lib"
    "github.com/gin-gonic/gin"
)

// Middleware to force a specific Content-Type header

type ContentTypeFilterMiddleware struct {
    logger  lib.Logger
}

func NewContentTypeFilterMiddleware(logger lib.Logger) ContentTypeFilterMiddleware {
    return ContentTypeFilterMiddleware{
        logger:  logger,
    }
}

func (this ContentTypeFilterMiddleware) Setup() {
    this.logger.Info("Setting up content type filter middleware")
}

func (this ContentTypeFilterMiddleware) Handler(allowedType string) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        if ctx.ContentType() != allowedType {
            ctx.AbortWithStatus(http.StatusUnsupportedMediaType)
            return
        }
        ctx.Next()
    }
}
