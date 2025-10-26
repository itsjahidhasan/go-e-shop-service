package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"go-e-shop-service/internal/infrastructure/db"
	apphttp "go-e-shop-service/internal/infrastructure/http"
	"go-e-shop-service/internal/infrastructure/http/handler"
	"go-e-shop-service/internal/repository"
	"go-e-shop-service/internal/usecase"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system environment")
	}

	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		log.Fatal("❌ POSTGRES_DSN not set in environment")
	}

	database, err := db.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	userRepo := repository.NewUserRepoPostgres(database)
	authUC := usecase.NewAuthUseCase(userRepo)
	authHandler := handler.NewAuthHandler(authUC)
	router := apphttp.SetupRouter(authHandler)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = ":8080"
	}

	log.Println("🚀 Server running on", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatal(err)
	}
}
