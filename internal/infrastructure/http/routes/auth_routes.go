package routes

import (
	"go-e-shop-service/internal/infrastructure/http/handler"
	"net/http"
)


func AuthRoutes(mux *http.ServeMux, authHandler *handler.AuthHandler) {
	mux.HandleFunc("/register", authHandler.Register)
}