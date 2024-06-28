package repository

import (
	"fmt"
	"strings"

	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	viewmodels "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/ViewModels"
	"github.com/jmoiron/sqlx"
)

type AssignPostgres struct {
	db *sqlx.DB
}

func NewAssignService(db *sqlx.DB) *AssignPostgres {
	return &AssignPostgres{db: db}
}

func (r *AssignPostgres) CreateAssign(assign models.Assign) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (assign_id, time_limit, memory_limit, correct_script, db_id, subtask_id, status_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", assignTable)
	row := r.db.QueryRow(query, assign.AssignID, assign.TimeLimit, assign.MemoryLimit, assign.CorrectScript, assign.DatabaseID, assign.SubtaskID, assign.StatusID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AssignPostgres) UpdateAssign(assign models.Assign) error {
	query := fmt.Sprintf("UPDATE %s SET time_limit=$1, memory_limit=$2, correct_script=$3, db_id=$4, status_id=$5 WHERE assign_id=$6 and subtask_id=$7", assignTable)
	_, err := r.db.Exec(query, assign.TimeLimit, assign.MemoryLimit, assign.CorrectScript, assign.DatabaseID, assign.StatusID, assign.AssignID, assign.SubtaskID)
	return err
}

func (r *AssignPostgres) GetAssignByID(id int) (viewmodels.AssignViewModel, error) {
	var assign viewmodels.AssignViewModel

	query := fmt.Sprintf("SELECT assign_id, time_limit, memory_limit, db_id FROM %s WHERE assign_id=$1 LIMIT 1", assignTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&assign.AssignID, &assign.TimeLimit, &assign.MemoryLimit, &assign.DatabaseID); err != nil {
		return assign, err
	}

	var correctScripts strings.Builder

	query = fmt.Sprintf("SELECT subtask_id, correct_script FROM %s WHERE assign_id=$1 ORDER BY subtask_id", assignTable)
	rows, err := r.db.Query(query, id)
	if err != nil {
		return assign, err
	}
	for rows.Next() {
		var key, value string
		if err := rows.Scan(&key, &value); err != nil {
			return assign, err
		}
		correctScripts.WriteString(key)
		correctScripts.WriteString(". ")
		correctScripts.WriteString(value)
		correctScripts.WriteString(";\n")
	}
	defer rows.Close()

	assign.CorrectScript = correctScripts.String()

	return assign, nil
}

func (r *AssignPostgres) GetAllAssignes() ([]models.Assign, error) {
	var assignes []models.Assign
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY assign_id, subtask_id", assignTable)
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

func (r *AssignPostgres) DeleteAssign(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE assign_id=$1", assignTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *AssignPostgres) GetAssignScriptsByID(id int) (map[int]models.Assign, error) {
	scripts := make(map[int]models.Assign)
	query := fmt.Sprintf("SELECT subtask_id, correct_script, status_id FROM %s WHERE assign_id = $1 ORDER BY subtask_id", assignTable)
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var temp models.Assign
		err := rows.Scan(&temp.SubtaskID, &temp.CorrectScript, &temp.StatusID)
		if err != nil {
			return scripts, err
		}
		scripts[temp.SubtaskID] = temp
	}
	return scripts, nil
}
