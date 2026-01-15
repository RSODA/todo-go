package todo

import "github.com/RSODA/todo-go/internal/service"

type Handler struct {
	s service.Service
}

func NewHandler(s service.Service) *Handler {
	return &Handler{s: s}
}
