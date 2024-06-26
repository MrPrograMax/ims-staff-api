package repository

import (
	"github.com/jmoiron/sqlx"
	"ims-staff-api/models"
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

type Repository struct {
	Task
	Status
	Staff
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Task:   NewTaskPostgres(db),
		Status: NewStatusPostgres(db),
		Staff:  NewStaffPostgres(db),
	}
}
