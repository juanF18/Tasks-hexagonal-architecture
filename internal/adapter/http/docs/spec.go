// internal/adapter/http/docs/spec.go
package docs

import (
	"net/http"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gin-gonic/gin"

	"test-hex-architecture/internal/adapter/http/docs/paths"
)

func RegisterDocsHandler(r *gin.Engine) {
	r.GET("/docs", func(c *gin.Context) {
		spec := buildOpenAPISpec()

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

func buildOpenAPISpec() map[string]interface{} {
	return map[string]interface{}{
		"openapi": "3.0.0",
		"info": map[string]interface{}{
			"title":       "Task API",
			"description": "A task management API with hexagonal architecture. **Tip**: Use pagination parameters (page & limit) for better performance.",
			"version":     "1.0.0",
		},
		"servers": []map[string]interface{}{
			{
				"url":         "http://localhost:8080",
				"description": "Development server",
			},
		},
		"paths":      buildPaths(),
		"components": buildComponents(),
		"x-samples": map[string]interface{}{ // Ejemplos globales
			"GET /tasks/?page=1&limit=10": "Recommended way to list tasks",
			"GET /tasks/":                 "Gets all tasks (use with caution)",
		},
	}
}

func buildPaths() map[string]interface{} {
	allPaths := make(map[string]interface{})

	// Agregar paths de tasks
	for path, spec := range paths.GetTaskPaths() {
		allPaths[path] = spec
	}

	// En el futuro: agregar paths de users, auth, etc.
	// for path, spec := range paths.GetUserPaths() {
	//     allPaths[path] = spec
	// }

	return allPaths
}
