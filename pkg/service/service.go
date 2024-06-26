package service

import (
	"ims-staff-api/models"
	"ims-staff-api/pkg/repository"
)

type Task interface {
	GetTasksList() ([]models.Task, error)
	GetTasksByStatus() ([]models.Task, error)
	GetTaskByID(id int64) (models.Task, error)
	CreateTask(task models.Task) (int64, error)
	UpdateTask(taskId int64, input models.Task) error
	DeleteTask() error
}

type Status interface {
	GetStatusList() ([]models.TaskStatus, error)
}

type Staff interface {
	//TODO GetStaffList()
	//TODO GetWorkerTasksById()
	//TODO AssignTask()
}

type Services struct {
	Task
	Status
	Staff
}

func NewService(repos *repository.Repository) *Services {
	return &Services{
		Task:   NewTaskService(repos.Task),
		Status: NewStatusService(repos.Status),
		Staff:  NewStaffService(repos.Staff),
	}
}
