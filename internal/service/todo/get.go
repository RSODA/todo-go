package todo

import (
	"context"
	"errors"

	"github.com/RSODA/todo-go/internal/models"
)

func (s *todoService) Get(ctx context.Context, id int64) (*models.TODO, error) {
	if id < 0 {
		return nil, errors.New("invalid id")
	}

	res, err := s.r.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
