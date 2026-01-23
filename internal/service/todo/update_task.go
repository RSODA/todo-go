package todo

import (
	"context"
	"errors"

	"github.com/RSODA/todo-go/internal/models"
)

func (s *todoService) UpdateTask(ctx context.Context, m *models.UpdateTaskRequest) error {
	if m == nil {
		return errors.New("UpdateTaskRequest cannot be nil")
	}

	err := s.r.UpdateTask(ctx, m)
	if err != nil {
		return err
	}

	return nil
}
