package repository

import (
	"fmt"

	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	"github.com/jmoiron/sqlx"
)

type StatusPostgres struct {
	db *sqlx.DB
}

func NewStatusService(db *sqlx.DB) *StatusPostgres {
	return &StatusPostgres{db: db}
}

func (r *StatusPostgres) GetAllStatuses() ([]models.Status, error) {
	var statuses []models.Status
	query := fmt.Sprintf("SELECT * FROM %s", statusTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tempStatus models.Status
		if err := rows.Scan(&tempStatus.ID, &tempStatus.Name); err != nil {
			return statuses, err
		}
		statuses = append(statuses, tempStatus)
	}
	if err = rows.Err(); err != nil {
		return statuses, err
	}
	return statuses, nil
}
