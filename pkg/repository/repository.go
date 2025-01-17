package repository

import (
	"github.com/jmoiron/sqlx"
	"ims-staff-api/models"
)

type Authorization interface {
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

type Repository struct {
	Task
	Status
	Staff
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Task:   NewTaskPostgres(db),
		Status: NewStatusPostgres(db),
		Staff:  NewStaffPostgres(db),
	}
}
