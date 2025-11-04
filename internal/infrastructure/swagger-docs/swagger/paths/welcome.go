package paths

var WelcomePath = map[string]interface{}{
	"/welcome": map[string]interface{}{
		"get": map[string]interface{}{
			"summary":     "Welcome message",
			"description": "Returns a simple welcome message for the E-Commerce API.",
			"tags":        []string{"General"},
			"responses": map[string]interface{}{
				"200": map[string]interface{}{
					"description": "Successful response with welcome message",
					"content": map[string]interface{}{
						"text/plain": map[string]interface{}{
							"schema": map[string]interface{}{
								"type":    "string",
								"example": "Welcome to the E-Commerce API!",
							},
						},
					},
				},
			},
		},
	},
}
