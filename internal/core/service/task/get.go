package taskservice

import (
	"context"
	"test-hex-architecture/internal/core/domain/task"
	"test-hex-architecture/internal/core/port"
	"test-hex-architecture/internal/shared/domain"
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

// Método unificado que maneja ambos casos
func (s *List) Execute(ctx context.Context, paginationParams *domain.PaginationParams) (interface{}, error) {
	if paginationParams == nil {
		tasks, err := s.Repo.FindAll(ctx, 0, 0) // Obtener todos los registros
		return tasks, err
	}

	tasks, err := s.Repo.FindAll(ctx, paginationParams.Offset(), paginationParams.Limit)

	if err != nil {
		return nil, err
	}

	totalItems, err := s.Repo.CountAll(ctx)
	if err != nil {
		return nil, err
	}

	return domain.NewPaginatedResponse(tasks, *paginationParams, int(totalItems)), nil
}
