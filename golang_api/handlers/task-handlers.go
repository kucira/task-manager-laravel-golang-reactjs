package handlers

import (
	"net/http"
	"task-manager-fullstack/models"
	"task-manager-fullstack/services"

	"github.com/gin-gonic/gin"
)

func GetDataTasks(c *gin.Context) {
	tasks, err := services.GetTasks()
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Something wrong", err)
		return
	}
	models.SuccessResponse(c, http.StatusOK, tasks)
}
