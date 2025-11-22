package services

import (
	"strconv"
	"task-manager-fullstack/config"
	"task-manager-fullstack/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) ([]models.Task, models.CursorPagination, error) {
	var tasks []models.Task
	pagination, scope := models.CursorPaginate(c)
	result := config.DB.Scopes(scope).Find(&tasks)
	// Calculate next cursor
	if len(tasks) > 0 {
		lastID := tasks[len(tasks)-1].ID
		pagination.NextCursor = strconv.Itoa(int(lastID))
		pagination.HasMore = len(tasks) == pagination.Limit
	}

	if result.Error != nil {
		return nil, pagination, result.Error
	}

	return tasks, pagination, nil
}
