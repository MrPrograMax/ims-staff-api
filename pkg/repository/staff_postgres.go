package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"ims-staff-api/models"
)

type StaffPostgres struct {
	db *sqlx.DB
}

func NewStaffPostgres(db *sqlx.DB) *StaffPostgres {
	return &StaffPostgres{db: db}
}

func (s *StaffPostgres) GetStaffList() ([]models.UserDto, error) {
	var workerList []models.UserDto
	query := fmt.Sprintf(`SELECT u.id, u.full_name, u.role_id FROM %s u`, userTable)
	err := s.db.Select(&workerList, query)

	return workerList, err
}

func (s *StaffPostgres) GetWorkerTasksById(workerId int64) ([]models.Task, error) {
	var workerTaskList []models.Task

	query := fmt.Sprintf(`SELECT * FROM %s t WHERE t.id = $1`, taskTable)
	err := s.db.Select(&workerTaskList, query, workerId)

	return workerTaskList, err
}

func (s *StaffPostgres) AssignTask() error {
	return nil
}
