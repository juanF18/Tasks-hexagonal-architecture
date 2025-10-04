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
		"description": "Get a list of all tasks",
		"tags":        []string{"tasks"},
		"responses": map[string]interface{}{
			"200": map[string]interface{}{
				"description": "List of tasks",
				"content": map[string]interface{}{
					"application/json": map[string]interface{}{
						"schema": map[string]interface{}{
							"type": "array",
							"items": map[string]interface{}{
								"$ref": "#/components/schemas/Task",
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
