package taskservice

import (
	"context"
	"errors"

	"test-hex-architecture/internal/core/port"
)

type Update struct{ Repo port.TaskRepository }

func NewUpdate(r port.TaskRepository) *Update { return &Update{Repo: r} }

func (s *Update) Execute(ctx context.Context, id, title string, description string, done *bool) error {
	t, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if t == nil {
		return errors.New("task not found")
	}
	if title != "" {
		if err := t.EditTask(title, description, *done); err != nil {
			return err
		}
	}
	if done != nil {
		t.MarkAsDoneOrUndone(*done)
	}
	return s.Repo.Update(ctx, t)
}
