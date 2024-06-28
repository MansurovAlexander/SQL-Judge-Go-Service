package repository

import (
	"database/sql"
	"fmt"

	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	"github.com/jmoiron/sqlx"
)

type SubmissionPostgres struct {
	db *sqlx.DB
}

func NewSubmissionService(db *sqlx.DB) *SubmissionPostgres {
	return &SubmissionPostgres{db: db}
}

func (r *SubmissionPostgres) CreateSubmission(submission models.Submission) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (submission_id, student_id, time, memory, script, status_id, assign_id, subtask_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", submissionTable)
	row := r.db.QueryRow(query, submission.SubmissionID, submission.StudentID, submission.Time, submission.Memory, submission.Script, submission.Status, submission.AssignID, submission.SubtaskID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *SubmissionPostgres) GetSubmissionByID(studentId, assignId int) ([]models.Submission, error) {
	var submissionList []models.Submission
	query := fmt.Sprintf("SELECT id, submission_id, student_id, time, memory, script, status_id, assign_id, subtask_id FROM %s WHERE student_id=$1 AND assign_id=$2 ORDER BY subtask_id", submissionTable)
	row, err := r.db.Query(query, studentId, assignId)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var tempModel models.Submission
		if err := row.Scan(&tempModel.ID, &tempModel.SubmissionID, &tempModel.StudentID, &tempModel.Time,
			&tempModel.Memory, &tempModel.Script, &tempModel.Status, &tempModel.AssignID, &tempModel.SubtaskID); err != nil {
			return nil, err
		}
		submissionList = append(submissionList, tempModel)
	}
	return submissionList, nil
}

func (r *SubmissionPostgres) GetAllSubmissions() ([]models.Submission, error) {
	var submissiones []models.Submission
	query := fmt.Sprintf("SELECT id, submission_id, student_id, time, memory, script, status_id, assign_id, subtask_id FROM %s", submissionTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var tempSubmission models.Submission
		if err := rows.Scan(&tempSubmission.ID, &tempSubmission.SubmissionID, &tempSubmission.StudentID, &tempSubmission.Time,
			&tempSubmission.Memory, &tempSubmission.Script, &tempSubmission.Status, &tempSubmission.AssignID, &tempSubmission.SubtaskID); err != nil {
			return submissiones, err
		}
		submissiones = append(submissiones, tempSubmission)
	}
	if err = rows.Err(); err != nil {
		return submissiones, err
	}
	return submissiones, nil
}

func (r *SubmissionPostgres) DeleteSubmissionsByAssignID(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE assign_id=$1", submissionTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *SubmissionPostgres) GetByUserAssignSubTaskID(assignID, subtaskID, studentID int) bool {
	query := fmt.Sprintf("SELECT 1 FROM %s WHERE assign_id = $1 AND subtask_id = $2 AND student_id = $3", submissionTable)
	rows := r.db.QueryRow(query, assignID, subtaskID, studentID)
	if err := rows.Scan(); err != nil {
		if err == sql.ErrNoRows {
			return false
		}
	}
	return true
}

func (r *SubmissionPostgres) UpdateSubmission(submission models.Submission) error {
	query := fmt.Sprintf("UPDATE %s SET script = $1, time = $2, memory = $3, status_id = $4 WHERE assign_id = $5 AND subtask_id = $6 AND student_id = $7", submissionTable)
	_, err := r.db.Exec(query, submission.Script, submission.Time, submission.Memory, submission.Status, submission.AssignID, submission.SubtaskID, submission.StudentID)
	if err != nil {
		return err
	}
	return nil
}
