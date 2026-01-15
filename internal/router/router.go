package router

import (
	"github.com/RSODA/todo-go/internal/handlers/todo"
	"github.com/gin-gonic/gin"
)

func NewRouter(h *todo.Handler) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery(), gin.Logger())

	r.POST("api/v1/todo/create", h.Create)

	return r
}
