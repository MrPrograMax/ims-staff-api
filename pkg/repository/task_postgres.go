package repository

import (
	"github.com/jmoiron/sqlx"
	"ims-staff-api/models"
)

type TaskPostgres struct {
	db *sqlx.DB
}

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}

func (t TaskPostgres) GetTasksList() ([]models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskPostgres) GetTasksByStatus() ([]models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskPostgres) GetTaskByID(id int64) (models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskPostgres) CreateTask(task models.Task) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (t TaskPostgres) UpdateTask(taskId int64, input models.Task) error {
	//TODO implement me
	panic("implement me")
}

func (t TaskPostgres) DeleteTask() error {
	//TODO implement me
	panic("implement me")
}
