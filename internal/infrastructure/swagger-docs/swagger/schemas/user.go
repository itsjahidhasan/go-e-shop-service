package schemas

var UserSchema = map[string]interface{}{
	"User": map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"id": map[string]interface{}{
				"type":    "string",
				"example": "64f1c2b7a9d4f23c45e2a9d3",
			},
			"name": map[string]interface{}{
				"type":    "string",
				"example": "John Doe",
			},
			"email": map[string]interface{}{
				"type":    "string",
				"example": "johndoe@gmail.com",
			},
			"phone": map[string]interface{}{
				"type":    "string",
				"example": "+1234567890",
			},
			"created_at": map[string]interface{}{
				"type":    "string",
				"format":  "date-time",
				"example": "2025-08-24T12:34:56Z",
			},
			"updated_at": map[string]interface{}{
				"type":    "string",
				"format":  "date-time",
				"example": "2025-08-25T08:22:33Z",
			},
		},
	},
}
