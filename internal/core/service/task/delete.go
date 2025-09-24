// internal/core/service/task/delete.go
package taskservice

import (
	"context"

	"test-hex-architecture/internal/core/port"
)

type Delete struct{ Repo port.TaskRepository }

func NewDelete(r port.TaskRepository) *Delete { return &Delete{Repo: r} }

func (s *Delete) Execute(ctx context.Context, id string) error {
	return s.Repo.Delete(ctx, id)
}
