package postgress

import (
	"github.com/RSODA/todo-go/internal/repo"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	tableTodo = "todo"
	tableTask = "task"

	idColumn          = "id"
	titleColumn       = "title"
	descriptionColumn = "description"
	isCompleteColumn  = "is_complete"
	createdAtColumn   = "created_at"
	todoIdColumn      = "todo_id"
)

type postgres struct {
	db *pgxpool.Pool
}

func NewPostgres(db *pgxpool.Pool) repo.Repo {
	return &postgres{
		db: db,
	}
}
