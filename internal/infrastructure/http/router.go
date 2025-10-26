// internal/infrastructure/http/router.go
package http

import (
	"go-e-shop-service/internal/infrastructure/http/handler"
	"net/http"
)

func SetupRouter(authHandler *handler.AuthHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/register", authHandler.Register)
	return mux
}
