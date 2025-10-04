// internal/adapter/http/docs/components.go
package docs

func buildComponents() map[string]interface{} {
	return map[string]interface{}{
		"schemas": map[string]interface{}{
			"Task":              getTaskSchema(),
			"CreateTaskRequest": getCreateTaskRequestSchema(),
			"UpdateTaskRequest": getUpdateTaskRequestSchema(),
			"ErrorResponse":     getErrorResponseSchema(),
			"IDResponse":        getIDResponseSchema(),
		},
	}
}

func getTaskSchema() map[string]interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"id": map[string]interface{}{
				"type":        "string",
				"description": "Task unique identifier",
				"example":     "123e4567-e89b-12d3-a456-426614174000",
			},
			"title": map[string]interface{}{
				"type":        "string",
				"description": "Task title",
				"example":     "Aprender hexagonal",
			},
			"description": map[string]interface{}{
				"type":        "string",
				"description": "Task description",
				"example":     "Hacer CRUD con Gin y Mongo v2",
			},
			"done": map[string]interface{}{
				"type":        "boolean",
				"description": "Task completion status",
				"example":     false,
			},
			"createdAt": map[string]interface{}{
				"type":        "string",
				"format":      "date-time",
				"description": "Task creation timestamp",
			},
			"updatedAt": map[string]interface{}{
				"type":        "string",
				"format":      "date-time",
				"description": "Task last update timestamp",
			},
		},
	}
}

func getCreateTaskRequestSchema() map[string]interface{} {
	return map[string]interface{}{
		"type":     "object",
		"required": []string{"title"},
		"properties": map[string]interface{}{
			"title": map[string]interface{}{
				"type":        "string",
				"description": "Task title",
				"example":     "Aprender hexagonal",
			},
			"description": map[string]interface{}{
				"type":        "string",
				"description": "Task description (optional)",
				"example":     "Hacer CRUD con Gin y Mongo v2",
			},
		},
	}
}

func getUpdateTaskRequestSchema() map[string]interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"title": map[string]interface{}{
				"type":        "string",
				"description": "Task title",
				"example":     "Hexagonal en producción",
			},
			"description": map[string]interface{}{
				"type":        "string",
				"description": "Task description",
				"example":     "CRUD completo con Gin y Mongo v2",
			},
			"done": map[string]interface{}{
				"type":        "boolean",
				"description": "Task completion status",
				"example":     true,
			},
		},
	}
}

func getErrorResponseSchema() map[string]interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"error": map[string]interface{}{
				"type":        "string",
				"description": "Error message",
				"example":     "Invalid payload",
			},
		},
	}
}

func getIDResponseSchema() map[string]interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"id": map[string]interface{}{
				"type":        "string",
				"description": "Created resource ID",
				"example":     "123e4567-e89b-12d3-a456-426614174000",
			},
		},
	}
}
