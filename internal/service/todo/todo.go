package todo

import (
	"github.com/RSODA/todo-go/internal/repo"
	"github.com/RSODA/todo-go/internal/service"
)

type todoService struct {
	r repo.Repo
}

func NewTODOService(r repo.Repo) service.Service {
	return &todoService{r: r}
}
