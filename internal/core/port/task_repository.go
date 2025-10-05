package port

import (
	"context"
	"test-hex-architecture/internal/core/domain/task"
)

type TaskRepository interface {
	Save(ctx context.Context, t *task.Task) (string, error)
	FindByID(ctx context.Context, id string) (*task.Task, error)
	FindAll(ctx context.Context, offset, limit int) ([]*task.Task, error)
	CountAll(ctx context.Context) (int64, error)
	Update(ctx context.Context, t *task.Task) error
	Delete(ctx context.Context, id string) error
}
