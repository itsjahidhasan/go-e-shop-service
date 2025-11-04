package swagger

import (
	"encoding/json"
	"go-e-shop-service/internal/config"
	"go-e-shop-service/internal/infrastructure/swagger-docs/swagger/paths"
	"go-e-shop-service/internal/infrastructure/swagger-docs/swagger/schemas"
	"go-e-shop-service/internal/infrastructure/swagger-docs/swagger/tags"
	"maps"
	"net/http"
	"sync"
)

var (
	cachedSwaggerJSON []byte
	cacheOnce         sync.Once
)

// mergeSwaggerElements merges multiple map[string]interface{} into one
func mergeSwaggerElements(el ...map[string]interface{}) map[string]interface{} {
	merged := make(map[string]interface{})
	for _, element := range el {
		maps.Copy(merged, element) // shallow copy
	}
	return merged
}

// SwaggerDocs serves the Swagger JSON
func PreGenerateSwagger() {
	// Generate & cache Swagger JSON only once
	cacheOnce.Do(func() {
		cfg := config.LoadConfig()

		sw := map[string]interface{}{
			"openapi": "3.0.0",
			"info": map[string]interface{}{
				"title":       "Go E-Shop API",
				"description": "This is an E-Commerce API service built with Go",
				"version":     "1.0.0",
			},
			"servers": []map[string]string{
				{"url": "http://localhost" + ensurePort(cfg.AppPort)},
			},
			"tags":       tags.Tags,
			"paths":      mergeSwaggerElements(paths.WelcomePath),
			"components": map[string]interface{}{"schemas": schemas.UserSchema},
		}

		var err error
		cachedSwaggerJSON, err = json.Marshal(sw)
		if err != nil {
			panic("Failed to generate Swagger JSON: " + err.Error())
		}
	})

}

func SwaggerDocs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(cachedSwaggerJSON)
}

// ensurePort guarantees the port string starts with ":"
func ensurePort(port string) string {
	if port == "" {
		return ":8080"
	}
	if port[0] != ':' {
		return ":" + port
	}
	return port
}
