package middlewares

import "go.uber.org/fx"

var Module = fx.Options(
    fx.Provide(NewCorsMiddleware),
    fx.Provide(NewMiddlewares),
    fx.Provide(NewContentTypeFilterMiddleware),
)

// Main middleware base package for dependency injection, define and setup all available middlewares

type IMiddleware interface {
    Setup()
}

type Middlewares []IMiddleware

func NewMiddlewares(
    corsMiddleware CorsMiddleware,
    contentTypeFilterMiddleware ContentTypeFilterMiddleware,
) Middlewares {
    return Middlewares{
        corsMiddleware,
        contentTypeFilterMiddleware,
    }
}

func (m Middlewares) Setup() {
    for _, middleware := range m {
        middleware.Setup()
    }
}
