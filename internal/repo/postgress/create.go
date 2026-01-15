package postgress

import (
	"context"

	sqr "github.com/Masterminds/squirrel"
	"github.com/RSODA/todo-go/internal/models"
)

func (p *postgres) Create(ctx context.Context, b *models.TODO) (*models.TODO, error) {
	var task models.Task

	builder := sqr.Insert(tableTodo).Columns(titleColumn, createdAtColumn).Values(b.Title, b.CreatedAt).Suffix("RETURNING *").PlaceholderFormat(sqr.Dollar)
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	err = p.db.QueryRow(ctx, query, args...).Scan(&b.ID, &b.Title, &b.CreatedAt)
	if err != nil {
		return nil, err
	}

	for _, v := range b.Tasks {
		builder = sqr.Insert(tableTask).Columns(taskIdColumn, descriptionColumn, isCompleteColumn).Values(b.ID, v.Description).Suffix("RETURNING *").PlaceholderFormat(sqr.Dollar)
		query, args, err = builder.ToSql()
		if err != nil {
			return nil, err
		}

		err = p.db.QueryRow(ctx, query, args...).Scan(&task.ID, &task.Description, &task.IsComplete)
		if err != nil {
			return nil, err
		}

		b.Tasks = append(b.Tasks, task)
	}

	return b, nil
}
