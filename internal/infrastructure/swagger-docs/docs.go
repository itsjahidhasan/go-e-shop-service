package swaggerdocs

import (
	"fmt"
	"go-e-shop-service/internal/infrastructure/swagger-docs/swagger"
	"net/http"
)

func SetSwaggerDocs(r *http.ServeMux) {
	r.HandleFunc("/swagger.json", swagger.SwaggerDocs)
	fs := http.FileServer(http.Dir("./internal/infrastructure/swagger-docs/swagger-ui"))
	r.Handle("/swagger/", http.StripPrefix("/swagger", fs))

	// 👇 Debug catch-all route (temporary)
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("REQ:", req.URL.Path)
		w.WriteHeader(http.StatusNotFound)
	})
}
