package repository

import (
	"fmt"
	"math/big"

	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	"github.com/jmoiron/sqlx"
)

type SubmissionPostgres struct {
	db *sqlx.DB
}

func NewSubmissionService(db *sqlx.DB) *SubmissionPostgres {
	return &SubmissionPostgres{db: db}
}

func (r *SubmissionPostgres) CreateSubmission(submission models.Submission) (big.Int, error) {
	var id big.Int
	query := fmt.Sprintf("INSERT INTO %s (student_id, time, memory, script, grade, assign_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", submissionTable)
	row := r.db.QueryRow(query, submission.StudentID, submission.Time, submission.Memory, submission.Script, submission.Grade, submission.AssignID)
	if err := row.Scan(&id); err != nil {
		return *big.NewInt(0), err
	}
	return id, nil
}

func (r *SubmissionPostgres) GetSubmissionByID(id big.Int) (models.Submission, error) {
	var submission models.Submission
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", submissionTable)
	row := r.db.QueryRow(query, id)
	if err := row.Scan(&submission.ID, &submission.StudentID, &submission.Time,
		&submission.Memory, &submission.Script, &submission.Grade, &submission.AssignID); err != nil {
		return submission, err
	}
	return submission, nil
}

func (r *SubmissionPostgres) GetAllSubmissions() ([]models.Submission, error) {
	var submissiones []models.Submission
	query := fmt.Sprintf("SELECT * FROM %s", submissionTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tempSubmission models.Submission
		if err := rows.Scan(&tempSubmission.ID, &tempSubmission.StudentID, &tempSubmission.Time,
			&tempSubmission.Memory, &tempSubmission.Script, &tempSubmission.Grade, &tempSubmission.AssignID); err != nil {
			return submissiones, err
		}
		submissiones = append(submissiones, tempSubmission)
	}
	if err = rows.Err(); err != nil {
		return submissiones, err
	}
	return submissiones, nil
}
