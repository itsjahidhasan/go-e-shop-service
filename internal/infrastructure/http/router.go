// internal/infrastructure/http/router.go
package http

import (
	"go-e-shop-service/internal/infrastructure/http/handler"
	"go-e-shop-service/internal/infrastructure/http/routes"
	"net/http"
)

func SetupRouter(authHandler *handler.AuthHandler) *http.ServeMux {
	mux := http.NewServeMux()
	routes.AuthRoutes(mux, authHandler)
	return mux
}
