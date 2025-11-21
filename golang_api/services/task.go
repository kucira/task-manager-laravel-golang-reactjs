package services

import (
	"task-manager-fullstack/config"
	"task-manager-fullstack/models"
)

func GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	result := config.DB.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}
