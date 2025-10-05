// internal/adapter/http/docs/paths/tasks.go
package paths

func GetTaskPaths() map[string]interface{} {
	return map[string]interface{}{
		"/tasks/":     getTaskCollectionPaths(),
		"/tasks/{id}": getTaskItemPaths(),
	}
}

func getTaskCollectionPaths() map[string]interface{} {
	return map[string]interface{}{
		"get":  getListTasksSpec(),
		"post": getCreateTaskSpec(),
	}
}

func getTaskItemPaths() map[string]interface{} {
	return map[string]interface{}{
		"get":    getGetTaskSpec(),
		"put":    getUpdateTaskSpec(),
		"delete": getDeleteTaskSpec(),
	}
}

func getListTasksSpec() map[string]interface{} {
	return map[string]interface{}{
		"summary":     "List all tasks",
		"description": "Get a list of all tasks with optional pagination. **Recommendation: Always use pagination parameters for better performance.**",
		"tags":        []string{"tasks"},
		"parameters": []map[string]interface{}{
			{
				"name":        "page",
				"in":          "query",
				"required":    false,
				"description": "Page number (starts from 1). **Default: 1 if limit is provided**",
				"schema": map[string]interface{}{
					"type":    "integer",
					"minimum": 1,
					"default": 1,
					"example": 1,
				},
			},
			{
				"name":        "limit",
				"in":          "query",
				"required":    false,
				"description": "Number of items per page (max 100). **Default: 10 if page is provided**",
				"schema": map[string]interface{}{
					"type":    "integer",
					"minimum": 1,
					"maximum": 100,
					"default": 10,
					"example": 10,
				},
			},
		},
		"responses": map[string]interface{}{
			"200": map[string]interface{}{
				"description": "List of tasks",
				"content": map[string]interface{}{
					"application/json": map[string]interface{}{
						"schema": map[string]interface{}{
							// Mostrar PRIMERO la respuesta paginada como ejemplo principal
							"oneOf": []map[string]interface{}{
								{
									"$ref":        "#/components/schemas/PaginatedTaskResponse",
									"description": "**Recommended**: Paginated tasks response (when page/limit provided)",
									"title":       "Paginated Response",
								},
								{
									"$ref":        "#/components/schemas/TaskArray",
									"description": "All tasks at once (when no pagination parameters). **Not recommended for production**",
									"title":       "All Tasks Array",
								},
							},
						},
						"examples": map[string]interface{}{
							"paginated": map[string]interface{}{
								"summary":     "Paginated Response (Recommended)",
								"description": "Example with pagination - GET /tasks/?page=1&limit=10",
								"value": map[string]interface{}{
									"data": []map[string]interface{}{
										{
											"id":          "123e4567-e89b-12d3-a456-426614174000",
											"title":       "Aprender hexagonal",
											"description": "Hacer CRUD con Gin y Mongo v2",
											"done":        false,
											"createdAt":   "2025-10-05T21:23:00Z",
											"updatedAt":   "2025-10-05T21:23:00Z",
										},
										{
											"id":          "456e7890-e89b-12d3-a456-426614174111",
											"title":       "Implementar paginación",
											"description": "Agregar paginación a la API",
											"done":        true,
											"createdAt":   "2025-10-05T20:15:00Z",
											"updatedAt":   "2025-10-05T21:20:00Z",
										},
									},
									"pagination": map[string]interface{}{
										"currentPage": 1,
										"pageSize":    10,
										"totalItems":  25,
										"totalPages":  3,
										"hasNext":     true,
										"hasPrev":     false,
									},
								},
							},
							"all_tasks": map[string]interface{}{
								"summary":     "All Tasks Array",
								"description": "Example without pagination - GET /tasks/ (not recommended for large datasets)",
								"value": []map[string]interface{}{
									{
										"id":          "123e4567-e89b-12d3-a456-426614174000",
										"title":       "Aprender hexagonal",
										"description": "Hacer CRUD con Gin y Mongo v2",
										"done":        false,
										"createdAt":   "2025-10-05T21:23:00Z",
										"updatedAt":   "2025-10-05T21:23:00Z",
									},
								},
							},
						},
					},
				},
			},
			"500": getErrorResponse("Internal server error"),
		},
	}
}

func getCreateTaskSpec() map[string]interface{} {
	return map[string]interface{}{
		"summary":     "Create a new task",
		"description": "Create a new task with title and description",
		"tags":        []string{"tasks"},
		"requestBody": map[string]interface{}{
			"required": true,
			"content": map[string]interface{}{
				"application/json": map[string]interface{}{
					"schema": map[string]interface{}{
						"$ref": "#/components/schemas/CreateTaskRequest",
					},
				},
			},
		},
		"responses": map[string]interface{}{
			"201": map[string]interface{}{
				"description": "Task created successfully",
				"content": map[string]interface{}{
					"application/json": map[string]interface{}{
						"schema": map[string]interface{}{
							"$ref": "#/components/schemas/IDResponse",
						},
					},
				},
			},
			"400": getErrorResponse("Invalid payload"),
			"422": getErrorResponse("Validation error"),
		},
	}
}

func getGetTaskSpec() map[string]interface{} {
	return map[string]interface{}{
		"summary":     "Get a task by ID",
		"description": "Get details of a specific task",
		"tags":        []string{"tasks"},
		"parameters":  []map[string]interface{}{getIDParameter()},
		"responses": map[string]interface{}{
			"200": map[string]interface{}{
				"description": "Task details",
				"content": map[string]interface{}{
					"application/json": map[string]interface{}{
						"schema": map[string]interface{}{
							"$ref": "#/components/schemas/Task",
						},
					},
				},
			},
			"404": getErrorResponse("Task not found"),
			"500": getErrorResponse("Internal server error"),
		},
	}
}

func getUpdateTaskSpec() map[string]interface{} {
	return map[string]interface{}{
		"summary":     "Update a task",
		"description": "Update an existing task completely",
		"tags":        []string{"tasks"},
		"parameters":  []map[string]interface{}{getIDParameter()},
		"requestBody": map[string]interface{}{
			"required": true,
			"content": map[string]interface{}{
				"application/json": map[string]interface{}{
					"schema": map[string]interface{}{
						"$ref": "#/components/schemas/UpdateTaskRequest",
					},
				},
			},
		},
		"responses": map[string]interface{}{
			"204": map[string]interface{}{
				"description": "Task updated successfully",
			},
			"400": getErrorResponse("Invalid payload"),
			"404": getErrorResponse("Task not found"),
			"422": getErrorResponse("Validation error"),
		},
	}
}

func getDeleteTaskSpec() map[string]interface{} {
	return map[string]interface{}{
		"summary":     "Delete a task",
		"description": "Delete a task by ID",
		"tags":        []string{"tasks"},
		"parameters":  []map[string]interface{}{getIDParameter()},
		"responses": map[string]interface{}{
			"204": map[string]interface{}{
				"description": "Task deleted successfully",
			},
			"500": getErrorResponse("Internal server error"),
		},
	}
}

// Helpers reutilizables
func getIDParameter() map[string]interface{} {
	return map[string]interface{}{
		"name":        "id",
		"in":          "path",
		"required":    true,
		"description": "Task ID",
		"schema": map[string]interface{}{
			"type": "string",
		},
	}
}

func getErrorResponse(description string) map[string]interface{} {
	return map[string]interface{}{
		"description": description,
		"content": map[string]interface{}{
			"application/json": map[string]interface{}{
				"schema": map[string]interface{}{
					"$ref": "#/components/schemas/ErrorResponse",
				},
			},
		},
	}
}
