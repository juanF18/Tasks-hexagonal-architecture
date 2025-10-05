// internal/adapter/http/docs/components.go
package docs

func buildComponents() map[string]interface{} {
	return map[string]interface{}{
		"schemas": map[string]interface{}{
			"Task":                  getTaskSchema(),
			"CreateTaskRequest":     getCreateTaskRequestSchema(),
			"UpdateTaskRequest":     getUpdateTaskRequestSchema(),
			"ErrorResponse":         getErrorResponseSchema(),
			"IDResponse":            getIDResponseSchema(),
			"PaginatedTaskResponse": getPaginatedTaskResponseSchema(), // nuevo
			"TaskArray":             getTaskArraySchema(),             // nuevo para compatibilidad
		},
	}
}

// Corregir: esto debe ser solo la entidad Task individual
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
				"example":     "2025-10-05T21:23:00Z",
			},
			"updatedAt": map[string]interface{}{
				"type":        "string",
				"format":      "date-time",
				"description": "Task last update timestamp",
				"example":     "2025-10-05T21:23:00Z",
			},
		},
	}
}

// Schema para respuesta paginada (estructura principal)
func getPaginatedTaskResponseSchema() map[string]interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"data": map[string]interface{}{
				"type": "array",
				"items": map[string]interface{}{
					"$ref": "#/components/schemas/Task",
				},
			},
			"pagination": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"currentPage": map[string]interface{}{
						"type":        "integer",
						"description": "Current page number",
						"example":     1,
					},
					"pageSize": map[string]interface{}{
						"type":        "integer",
						"description": "Number of items per page",
						"example":     10,
					},
					"totalItems": map[string]interface{}{
						"type":        "integer",
						"format":      "int64",
						"description": "Total number of items in database",
						"example":     25,
					},
					"totalPages": map[string]interface{}{
						"type":        "integer",
						"description": "Total number of pages",
						"example":     3,
					},
					"hasNext": map[string]interface{}{
						"type":        "boolean",
						"description": "Whether there are more pages after current",
						"example":     true,
					},
					"hasPrev": map[string]interface{}{
						"type":        "boolean",
						"description": "Whether there are pages before current",
						"example":     false,
					},
				},
			},
		},
	}
}

// Schema para array simple de tasks (cuando no hay paginación)
func getTaskArraySchema() map[string]interface{} {
	return map[string]interface{}{
		"type": "array",
		"items": map[string]interface{}{
			"$ref": "#/components/schemas/Task",
		},
	}
}

// Los demás schemas permanecen igual...
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
