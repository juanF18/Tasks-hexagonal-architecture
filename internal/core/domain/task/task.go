package task

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          string
	Title       string
	Description string
	Done        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type PaginationParams struct {
	Page  int
	Limit int
}

func NewTask(title, description string) (*Task, error) {
	if title == "" {
		return nil, errors.New("invalid task parameters")
	}
	return &Task{
		ID:          uuid.NewString(),
		Title:       title,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (t *Task) EditTask(title, description string, done bool) error {
	if title == "" {
		return errors.New("invalid task parameters")
	}
	t.Title = title
	t.Description = description
	t.Done = done
	t.UpdatedAt = time.Now()
	return nil
}

func (t *Task) MarkAsDoneOrUndone(done bool) {
	t.Done = done
	t.UpdatedAt = time.Now()
}
