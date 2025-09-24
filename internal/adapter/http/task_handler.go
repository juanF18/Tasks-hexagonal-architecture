// internal/adapter/http/task_handler.go
package httpadapter

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	taskservice "test-hex-architecture/internal/core/service/task"
)

type TaskHandler struct {
	Create *taskservice.Create
	Get    *taskservice.GetByID
	List   *taskservice.List
	Update *taskservice.Update
	Delete *taskservice.Delete
}

func NewTaskHandler(c *taskservice.Create, g *taskservice.GetByID, l *taskservice.List, u *taskservice.Update, d *taskservice.Delete) *TaskHandler {
	return &TaskHandler{Create: c, Get: g, List: l, Update: u, Delete: d}
}

func (h *TaskHandler) Register(r *gin.Engine) {
	group := r.Group("tasks")
	group.POST("/", h.create)
	group.GET("/:id", h.getByID)
	group.GET("/", h.list)
	group.PUT("/:id", h.update)
	group.DELETE("/:id", h.delete)
}

type createReq struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type updateReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        *bool  `json:"done"`
}

func (h *TaskHandler) create(c *gin.Context) {
	var req createReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	id, err := h.Create.Execute(ctx, req.Title, req.Description)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (h *TaskHandler) getByID(c *gin.Context) {
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	t, err := h.Get.Execute(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch"})
		return
	}
	if t == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, t)
}

func (h *TaskHandler) list(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	items, err := h.List.Execute(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list"})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *TaskHandler) update(c *gin.Context) {
	id := c.Param("id")
	var req updateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := h.Update.Execute(ctx, id, req.Title, req.Description, req.Done); err != nil {
		if err.Error() == "task not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *TaskHandler) delete(c *gin.Context) {
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := h.Delete.Execute(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete"})
		return
	}
	c.Status(http.StatusNoContent)
}
