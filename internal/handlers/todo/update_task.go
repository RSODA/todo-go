package todo

import (
	"net/http"

	"github.com/RSODA/todo-go/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateTask(c *gin.Context) {
	var req *models.UpdateTaskRequest

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.s.UpdateTask(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task updated"})
}
