package postgres

import (
	"context"

	sqr "github.com/Masterminds/squirrel"
)

func (p *postgres) Delete(ctx context.Context, id int64) error {
	builder := sqr.Delete(tableTodo).Where(sqr.Eq{idColumn: id}).PlaceholderFormat(sqr.Dollar)
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
