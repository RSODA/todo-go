package todo

import (
	"context"
	"errors"

	"github.com/RSODA/todo-go/internal/models"
)

func (s *todoService) Create(ctx context.Context, b *models.TODO) (*models.TODO, error) {

	if len(b.Title) == 0 || len(b.Tasks) == 0 {
		return nil, errors.New("title or tasks required")
	}

	res, err := s.r.Create(ctx, b)

	if err != nil {
		return nil, err
	}

	return res, nil
}
