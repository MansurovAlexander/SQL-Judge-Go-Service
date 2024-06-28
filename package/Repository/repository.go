package repository

import (
	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	viewmodels "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/ViewModels"
	"github.com/jmoiron/sqlx"
)

type Database interface {
	DeleteDataBaseByID(id int) (string, error)
	CreateDatabase(database models.Database) (int, error)
	GetAllDatabases() ([]viewmodels.DatabaseViewModel, error)
	GetDatabaseByID(id int) (viewmodels.DatabaseViewModel, error)
}

type Dbms interface {
	CreateDbms(dbms models.Dbms) (int, error)
	GetAllDbms() ([]models.Dbms, error)
	GetDbmsByID(id int) (models.Dbms, error)
}

type Assign interface {
	CreateAssign(assign models.Assign) (int, error)
	UpdateAssign(assign models.Assign) error
	GetAssignByID(id int) (viewmodels.AssignViewModel, error)
	GetAllAssignes() ([]models.Assign, error)
	DeleteAssign(id int) error
	GetAssignScriptsByID(id int) (map[int]models.Assign, error)
}

type BannedWordToAssign interface {
	CreateBannedWordToAssign(assignID int, bannedWords, admissionWords map[string][]string) error
	GetBannedWordByAssignID(assignId int) (map[string]string, error)
	GetAdmissionWordByAssignID(id int) (map[string]string, error)
	DropAllBannedWordsByAssignID(assignId int) error
}

type Submission interface {
	CreateSubmission(submission models.Submission) (int, error)
	GetSubmissionByID(student_id, assign_id int) ([]models.Submission, error)
	GetAllSubmissions() ([]models.Submission, error)
	DeleteSubmissionsByAssignID(id int) error
	GetByUserAssignSubTaskID(assignID, subtaskID, studentID int) bool
	UpdateSubmission(submission models.Submission) error
}

type Status interface {
	GetAllStatuses() ([]models.Status, error)
}

type Repository struct {
	Database
	Dbms
	Assign
	BannedWordToAssign
	Submission
	Status
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Dbms:               NewDbmsService(db),
		Database:           NewDatabaseService(db),
		Assign:             NewAssignService(db),
		BannedWordToAssign: NewBannedWordToAssignService(db),
		Submission:         NewSubmissionService(db),
		Status:             NewStatusService(db),
	}
}
