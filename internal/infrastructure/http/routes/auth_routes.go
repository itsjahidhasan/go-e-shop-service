package routes

import (
	"database/sql"
	"go-e-shop-service/internal/infrastructure/http/handler"
	"go-e-shop-service/internal/repository"
	"go-e-shop-service/internal/usecase"
	"net/http"
)


func AuthRoutes(mux *http.ServeMux, db *sql.DB) {
	userRepo := repository.NewUserRepoPostgres(db)
	authUC := usecase.NewAuthUseCase(userRepo)
	authHandler := handler.NewAuthHandler(authUC)

	//this are handler routes
	mux.HandleFunc("/register", authHandler.Register)
}