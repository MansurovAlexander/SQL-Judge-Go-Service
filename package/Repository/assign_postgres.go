package repository

import (
	"fmt"
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
	query := fmt.Sprintf("INSERT INTO %s (id, time_limit, memory_limit, correct_script, db_id) VALUES ($1, $2, $3, $4, $5) RETURNING id", assignTable)
	row := r.db.QueryRow(query, assign.ID, assign.TimeLimit, assign.MemoryLimit, assign.CorrectScript, assign.DatabaseID)
	if err := row.Scan(&id); err != nil {
		return *big.NewInt(0), err
	}
	return id, nil
}

func (r *AssignPostgres) GetAssignByID(id big.Int) (models.Assign, error) {
	var assign models.Assign
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", assignTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&assign.ID, &assign.TimeLimit, &assign.MemoryLimit, &assign.CorrectScript, &assign.DatabaseID); err != nil {
		return assign, err
	}
	return assign, nil
}

func (r *AssignPostgres) GetAllAssignes() ([]models.Assign, error) {
	var assignes []models.Assign
	query := fmt.Sprintf("SELECT * FROM %s", assignTable)
	rows, err := r.db.Query(query)
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
