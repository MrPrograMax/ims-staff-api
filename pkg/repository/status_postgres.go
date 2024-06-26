package repository

import (
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

func (r StatusPostgres) GetStatusList() ([]models.TaskStatus, error) {
	var statusList []models.TaskStatus

	query := fmt.Sprintf("SELECT * FROM taskStatus")
	err := r.db.Select(&statusList, query)

	return statusList, err
}
