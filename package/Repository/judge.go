package repository

import (
	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	viewmodels "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/ViewModels"
	"github.com/jmoiron/sqlx"
)

type JudgePostgres struct {
	db *sqlx.DB
}

func NewJudgeService(db *sqlx.DB) *JudgePostgres {
	return &JudgePostgres{db: db}
}

func (r *JudgePostgres) CheckSubmission(submission models.Submission) (string, error) {
	transaction, err := r.db.Begin()
	if err != nil {
		return "", err
	}
	var (
		assign           viewmodels.AssignViewModel
		dbCreationScript string
		dbmsName         string
		//bannedWords []string
	)
	query := "(SELECT a.id, a.time_limit, a.memory_limit, a.correct_script, db.creation_script, dbms.name FROM assign a JOIN databases db on (a.db_id=db.id)) JOIN dbms on (db.dbms_id=dbms.id) WHERE a.id=$1"
	row := r.db.QueryRow(query, submission.AssignID)
	if err := row.Scan(&assign.ID, &assign.TimeLimit, &assign.MemoryLimit, &assign.CorrectScript, &dbCreationScript, &dbmsName); err != nil {
		return "", err
	}
	if err := transaction.Rollback(); err != nil {
		return "", err
	}
	return "test", nil
}
