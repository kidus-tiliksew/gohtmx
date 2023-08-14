package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kidus-tiliksew/gohtmx/repository"
)

type TodoHandlers struct {
	TodoRepository repository.TodoRepository
}

func (t *TodoHandlers) Index(c *gin.Context) {
	todos, err := t.TodoRepository.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.HTML(http.StatusOK, "index.html", todos)
}
