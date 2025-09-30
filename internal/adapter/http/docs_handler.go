package httpadapter

import (
	"net/http"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gin-gonic/gin"
)

func RegisterDocsHandler(r *gin.Engine) {
	r.GET("/docs", func(c *gin.Context) {
		// Spec OpenAPI creado manualmente o desde estructura
		spec := map[string]interface{}{
			"openapi": "3.0.0",
			"info": map[string]interface{}{
				"title":       "Task API",
				"description": "A task management API with hexagonal architecture",
				"version":     "1.0.0",
			},
			"servers": []map[string]interface{}{
				{
					"url":         "http://localhost:8080",
					"description": "Development server",
				},
			},
			"paths": map[string]interface{}{
				"/tasks/": map[string]interface{}{
					"get": map[string]interface{}{
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
						},
					},
					"post": map[string]interface{}{
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
											"type": "object",
											"properties": map[string]interface{}{
												"id": map[string]interface{}{
													"type":    "string",
													"example": "123e4567-e89b-12d3-a456-426614174000",
												},
											},
										},
									},
								},
							},
						},
					},
				},
				"/tasks/{id}": map[string]interface{}{
					"get": map[string]interface{}{
						"summary":     "Get a task by ID",
						"description": "Get details of a specific task",
						"tags":        []string{"tasks"},
						"parameters": []map[string]interface{}{
							{
								"name":        "id",
								"in":          "path",
								"required":    true,
								"description": "Task ID",
								"schema": map[string]interface{}{
									"type": "string",
								},
							},
						},
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
							"404": map[string]interface{}{
								"description": "Task not found",
							},
						},
					},
					"put": map[string]interface{}{
						"summary":     "Update a task",
						"description": "Update an existing task completely",
						"tags":        []string{"tasks"},
						"parameters": []map[string]interface{}{
							{
								"name":        "id",
								"in":          "path",
								"required":    true,
								"description": "Task ID",
								"schema": map[string]interface{}{
									"type": "string",
								},
							},
						},
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
							"404": map[string]interface{}{
								"description": "Task not found",
							},
						},
					},
					"delete": map[string]interface{}{
						"summary":     "Delete a task",
						"description": "Delete a task by ID",
						"tags":        []string{"tasks"},
						"parameters": []map[string]interface{}{
							{
								"name":        "id",
								"in":          "path",
								"required":    true,
								"description": "Task ID",
								"schema": map[string]interface{}{
									"type": "string",
								},
							},
						},
						"responses": map[string]interface{}{
							"204": map[string]interface{}{
								"description": "Task deleted successfully",
							},
						},
					},
				},
			},
			"components": map[string]interface{}{
				"schemas": map[string]interface{}{
					"Task": map[string]interface{}{
						"type": "object",
						"properties": map[string]interface{}{
							"id": map[string]interface{}{
								"type":    "string",
								"example": "123e4567-e89b-12d3-a456-426614174000",
							},
							"title": map[string]interface{}{
								"type":    "string",
								"example": "Aprender hexagonal",
							},
							"description": map[string]interface{}{
								"type":    "string",
								"example": "Hacer CRUD con Gin y Mongo v2",
							},
							"done": map[string]interface{}{
								"type":    "boolean",
								"example": false,
							},
							"createdAt": map[string]interface{}{
								"type":   "string",
								"format": "date-time",
							},
							"updatedAt": map[string]interface{}{
								"type":   "string",
								"format": "date-time",
							},
						},
					},
					"CreateTaskRequest": map[string]interface{}{
						"type":     "object",
						"required": []string{"title"},
						"properties": map[string]interface{}{
							"title": map[string]interface{}{
								"type":    "string",
								"example": "Aprender hexagonal",
							},
							"description": map[string]interface{}{
								"type":    "string",
								"example": "Hacer CRUD con Gin y Mongo v2",
							},
						},
					},
					"UpdateTaskRequest": map[string]interface{}{
						"type": "object",
						"properties": map[string]interface{}{
							"title": map[string]interface{}{
								"type":    "string",
								"example": "Hexagonal en producción",
							},
							"description": map[string]interface{}{
								"type":    "string",
								"example": "CRUD completo con Gin y Mongo v2",
							},
							"done": map[string]interface{}{
								"type":    "boolean",
								"example": true,
							},
						},
					},
				},
			},
		}

		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecContent: spec,
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Task API Documentation",
			},
			DarkMode: true,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate docs"})
			return
		}

		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, htmlContent)
	})
}
