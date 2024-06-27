package service

import (
	"ims-staff-api/models"
	"ims-staff-api/pkg/repository"
)

type TaskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *TaskService {
	return &TaskService{repo: repo}
}

func (t *TaskService) GetTasksList() ([]models.Task, error) {
	return t.repo.GetTasksList()
}

func (t *TaskService) GetTasksByStatus(statusName string) ([]models.Task, error) {
	return t.repo.GetTasksByStatus(statusName)
}

func (t *TaskService) GetTaskByID(id int64) (models.Task, error) {
	return t.repo.GetTaskByID(id)
}

func (t *TaskService) CreateTask(task models.Task, statusId int64) (int64, error) {
	return t.repo.CreateTask(task, statusId)
}

func (t *TaskService) UpdateTask(taskId int64, input models.Task) error {
	return t.repo.UpdateTask(taskId, input)
}

func (t *TaskService) DeleteTask(taskId, statusId int64) error {
	return t.repo.DeleteTask(taskId, statusId)
}
