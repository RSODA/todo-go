package postgres

import (
	"context"

	sqr "github.com/Masterminds/squirrel"
	"github.com/RSODA/todo-go/internal/models"
)

func (p *postgres) Get(ctx context.Context, id int64) (*models.TODO, error) {
	var todo models.TODO

	builder := sqr.Select(idColumn, titleColumn, createdAtColumn).From(tableTodo).Where(sqr.Eq{idColumn: id}).PlaceholderFormat(sqr.Dollar)
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	err = p.db.QueryRow(ctx, query, args...).Scan(&todo.ID, &todo.Title, &todo.CreatedAt)
	if err != nil {
		return nil, err
	}

	builder = sqr.Select(idColumn, todoIdColumn, descriptionColumn, isCompleteColumn).From(tableTask).Where(sqr.Eq{todoIdColumn: todo.ID}).PlaceholderFormat(sqr.Dollar)
	query, args, err = builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var task models.Task

		err = rows.Scan(&task.ID, &task.TodoID, &task.Description, &task.IsComplete)
		if err != nil {
			return nil, err
		}

		todo.Tasks = append(todo.Tasks, task)
	}

	return &todo, nil
}
