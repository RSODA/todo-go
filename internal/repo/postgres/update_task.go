package postgres

import (
	"context"

	sqr "github.com/Masterminds/squirrel"
	"github.com/RSODA/todo-go/internal/models"
)

func (p *postgres) UpdateTask(ctx context.Context, m *models.UpdateTaskRequest) error {
	builder := sqr.Update(tableTask).Where(sqr.Eq{idColumn: m.ID}).PlaceholderFormat(sqr.Dollar).Set(isCompleteColumn, m.IsComplete)
	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	_, err = p.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
