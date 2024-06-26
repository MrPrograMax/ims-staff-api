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

func (t TaskService) GetTasksList() ([]models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskService) GetTasksByStatus() ([]models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskService) GetTaskByID(id int64) (models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskService) CreateTask(task models.Task) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskService) UpdateTask(taskId int64, input models.Task) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskService) DeleteTask() error {
	//TODO implement me
	panic("implement me")
}
