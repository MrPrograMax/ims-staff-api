package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"ims-staff-api/models"
)

type StatusPostgres struct {
	db *sqlx.DB
}

func NewStatusPostgres(db *sqlx.DB) *StatusPostgres {
	return &StatusPostgres{db: db}
}

func (r *StatusPostgres) GetStatusList() ([]models.TaskStatus, error) {
	var statusList []models.TaskStatus

	query := fmt.Sprintf("SELECT * FROM %s", taskStatusTable)
	err := r.db.Select(&statusList, query)

	return statusList, err
}

func (r *StatusPostgres) GetStatusIdByName(status string) (int64, error) {
	var statusId int64
	query := fmt.Sprintf("SELECT id FROM %s WHERE name=$1", taskStatusTable)
	err := r.db.Get(&statusId, query, status)

	if statusId == 0 {
		return -1, errors.New("Wait status for task not found")
	}

	return statusId, err
}
