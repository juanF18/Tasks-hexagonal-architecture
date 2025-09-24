package taskservice

import (
	"context"
	"errors"
	"test-hex-architecture/internal/core/domain/task"
	"test-hex-architecture/internal/core/port"
)

type Create struct{ Repo port.TaskRepository }

func NewCreate(repo port.TaskRepository) *Create {
	return &Create{Repo: repo}
}

func (s *Create) Execute(ctx context.Context, title string, description string) (string, error) {
	if title == "" {
		return "", errors.New("title is required")
	}

	t, err := task.NewTask(title, description)

	if err != nil {
		return "", err
	}

	return s.Repo.Save(ctx, t)
}
