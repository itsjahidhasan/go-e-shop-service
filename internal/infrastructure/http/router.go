// internal/infrastructure/http/router.go
package http

import (
	"database/sql"
	"go-e-shop-service/internal/infrastructure/http/routes"
	swaggerdocs "go-e-shop-service/internal/infrastructure/swagger-docs"
	"net/http"
)

func SetupRouter(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	swaggerdocs.SetSwaggerDocs(mux)
	// welcome routes
	mux.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the E-Commerce API!"))
	})

	routes.AuthRoutes(mux, db)
	return mux
}
