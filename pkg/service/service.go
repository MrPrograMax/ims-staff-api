package service

import (
	"ims-staff-api/models"
	"ims-staff-api/pkg/repository"
)

type Authorization interface {
	ParseToken(token string) (int64, error)
}

type Task interface {
	GetTasksList() ([]models.Task, error)
	GetTasksByStatus(statusName string) ([]models.Task, error)
	GetTaskByID(id int64) (models.Task, error)
	CreateTask(task models.Task, statusId int64) (int64, error)
	UpdateTask(taskId int64, input models.UpdateTask) error
	DeleteTask() error
}

type Status interface {
	GetStatusList() ([]models.TaskStatus, error)
	GetStatusIdByName(status string) (int64, error)
}

type Staff interface {
	GetStaffList() ([]models.UserDto, error)
	GetWorkerTasksById(workerId int64) ([]models.Task, error)
	AssignTask() error
}

type Services struct {
	Task
	Status
	Staff
	Authorization
}

func NewService(repos *repository.Repository) *Services {
	return &Services{
		Task:   NewTaskService(repos.Task),
		Status: NewStatusService(repos.Status),
		Staff:  NewStaffService(repos.Staff),
	}
}
