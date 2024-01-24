package repository

import (
	"math/big"

	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	"github.com/jmoiron/sqlx"
)

type AssignPostgres struct {
	db *sqlx.DB
}

func NewAssignService(db *sqlx.DB) *AssignPostgres {
	return &AssignPostgres{db: db}
}

func (r *AssignPostgres) CreateAssign(assign models.Assign) (big.Int, error) {
	var id big.Int
	query := "INSERT INTO $1 (id, time_limit, memory_limit, correct_script, db_id) VALUES ($2, $3, $4, $5, $6) RETURNING id"
	row := r.db.QueryRow(query, assignTable, assign.ID, assign.TimeLimit, assign.MemoryLimit, assign.CorrectScript, assign.DatabaseID)
	if err := row.Scan(&id); err != nil {
		return *big.NewInt(0), err
	}
	return id, nil
}

func (r *AssignPostgres) GetAssignByID(id big.Int) (models.Assign, error) {
	var assign models.Assign
	query := "SELECT * FROM $1 WHERE id=$2"
	row := r.db.QueryRow(query, assignTable, id)
	if err := row.Scan(&assign.ID, &assign.TimeLimit, &assign.MemoryLimit, &assign.CorrectScript, &assign.DatabaseID); err != nil {
		return assign, err
	}
	return assign, nil
}

func (r *AssignPostgres) GetAllAssignes() ([]models.Assign, error) {
	var assignes []models.Assign
	query := "SELECT * FROM $1"
	rows, err := r.db.Query(query, assignTable)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tempAssign models.Assign
		if err := rows.Scan(&tempAssign.ID, &tempAssign.TimeLimit, &tempAssign.MemoryLimit, &tempAssign.CorrectScript,
			&tempAssign.DatabaseID); err != nil {
			return assignes, err
		}
		assignes = append(assignes, tempAssign)
	}
	if err = rows.Err(); err != nil {
		return assignes, err
	}
	return assignes, nil
}
