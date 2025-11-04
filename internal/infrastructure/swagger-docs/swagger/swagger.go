package swagger

import (
	"encoding/json"
	"go-e-shop-service/internal/config"
	"go-e-shop-service/internal/infrastructure/swagger-docs/swagger/schemas"
	"maps"
	"net/http"
)

func mergeSwaggerElements(el ...map[string]interface{}) map[string]interface{} {
	merged := make(map[string]interface{})
	for _, element := range el {
		maps.Copy(merged, element)
	}
	return merged
}

func SwaggerDocs(w http.ResponseWriter, r *http.Request) {
	cfg := config.LoadConfig()

	sw := map[string]interface{}{
		"openapi": "3.0.0",
		"info": map[string]interface{}{
			"title":       "Go E-Commerce API",
			"description": "This is a E-Commerce API service built with Go + Chi",
			"version":     "1.0.0",
		},
		"servers": []map[string]string{
			{
				"url": "http://localhost:" + cfg.AppPort,
			},
		},
		"components": map[string]interface{}{
			"schemas": map[string]interface{}{
				"User": schemas.UserSchema["User"],
			},
		},
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(sw)
}
