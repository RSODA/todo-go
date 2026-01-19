package todo

import (
	"net/http"

	"github.com/RSODA/todo-go/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	var todo models.TODO

	if err := c.ShouldBindBodyWithJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.s.Create(c, &todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	
	c.JSON(http.StatusOK, res)
}
