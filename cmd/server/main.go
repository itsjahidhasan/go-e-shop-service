package main

import (
	"log"
	"net/http"

	"go-e-shop-service/internal/config"
	"go-e-shop-service/internal/infrastructure/db"
	apphttp "go-e-shop-service/internal/infrastructure/http"
)

func main() {
	// Load environment variables
	cfg := config.LoadConfig()

	dsn := cfg.PostGresConnString
	if dsn == "" {
		log.Fatal("❌ POSTGRES_DSN not set in environment")
	}
	log.Println("✅ DB Path:", dsn)
	database, err := db.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	
	router := apphttp.SetupRouter(database)



	log.Println("🚀 Server running on:", cfg.AppServer)
	if err := http.ListenAndServe(cfg.AppPort, router); err != nil {
		log.Fatal("error starting server:", err)
	}
}
