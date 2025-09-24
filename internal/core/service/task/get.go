package taskservice

import (
	"context"
	"test-hex-architecture/internal/core/domain/task"
	"test-hex-architecture/internal/core/port"
)

type GetByID struct{ Repo port.TaskRepository }
type List struct{ Repo port.TaskRepository }

func NewGetByID(repo port.TaskRepository) *GetByID {
	return &GetByID{Repo: repo}
}

func NewList(repo port.TaskRepository) *List {
	return &List{Repo: repo}
}

func (s *GetByID) Execute(ctx context.Context, id string) (*task.Task, error) {
	return s.Repo.FindByID(ctx, id)
}

func (s *List) Execute(ctx context.Context) ([]*task.Task, error) {
	return s.Repo.FindAll(ctx)
}
