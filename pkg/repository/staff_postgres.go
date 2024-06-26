package repository

import "github.com/jmoiron/sqlx"

type StaffPostgres struct {
	db *sqlx.DB
}

func NewStaffPostgres(db *sqlx.DB) *StaffPostgres {
	return &StaffPostgres{db: db}
}
