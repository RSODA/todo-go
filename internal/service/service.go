package service

import (
	"context"

	"github.com/RSODA/todo-go/internal/models"
)

type Service interface {
	Create(ctx context.Context, b *models.TODO) (*models.TODO, error)
	Get(ctx context.Context, id int64) (*models.TODO, error)
}
