package repository

import (
	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	"github.com/jmoiron/sqlx"
)

type Database interface {
	CreateDatabase(database models.Database) (int, error)
	GetAllDatabases() ([]models.Database, error)
	GetDatabaseByID(id int) (models.Database, error)
}

type Dbms interface {
	CreateDbms(dbms models.Dbms) (int, error)
	GetAllDbms() ([]models.Dbms, error)
	GetDbmsByID(id int) (models.Dbms, error)
}

type BannedWord interface {
	CreateBannedWord(bannedWord models.BannedWord) (int, error)
	GetBannedWordByID(id int) (models.BannedWord, error)
	GetAllBannedWords() ([]models.BannedWord, error)
}

type Assign interface {
	CreateAssign(assign models.Assign) (int, error)
	GetAssignByID(id int) (models.Assign, error)
	GetAllAssignes() ([]models.Assign, error)
}

type BannedWordToAssign interface {
	CreateBannedWordToAssign(assignId int, bannedWordID int) (int, error)
	GetBannedWordByAssignID(assignId int) ([]int, error)
}

type Submission interface {
	CreateSubmission(submission models.Submission) (int, error)
	GetSubmissionByID(id int) (models.Submission, error)
	GetAllSubmissions() ([]models.Submission, error)
}

type Judge interface {
	CheckSubmission(inputedScript, dbCreateScript, correctScript, dbName string, assignID, studentID int, timeLimit, memoryLimit int) (int, error)
}

type Repository struct {
	Database
	Dbms
	Assign
	BannedWord
	BannedWordToAssign
	Submission
	Judge
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Dbms:               NewDbmsService(db),
		Database:           NewDatabaseService(db),
		Assign:             NewAssignService(db),
		BannedWord:         NewBannedWordService(db),
		BannedWordToAssign: NewBannedWordToAssignService(db),
		Submission:         NewSubmissionService(db),
		Judge:              NewJudgeService(db),
	}
}
