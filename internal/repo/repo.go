package repo

import (
	"context"

	"github.com/RSODA/todo-go/internal/models"
)

type Repo interface {
	Create(ctx context.Context, b *models.TODO) (*models.TODO, error)
}
