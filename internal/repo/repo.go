package repo

import (
	"context"

	"github.com/RSODA/todo-go/internal/models"
)

type Repo interface {
	Get(ctx context.Context, id int64) (*models.TODO, error)
	Create(ctx context.Context, b *models.TODO) (*models.TODO, error)
}
