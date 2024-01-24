package repository

import (
	"math/big"

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
	CreateAssign(assign models.Assign) (big.Int, error)
	GetAssignByID(id big.Int) (models.Assign, error)
	GetAllAssignes() ([]models.Assign, error)
}

type BannedWordToAssign interface {
	CreateBannedWordToAssign(models.BannedWordToAssign) (big.Int, error)
	GetBannedWordByAssignID(assignId big.Int) ([]int, error)
}

type Submission interface {
	CreateSubmission(submission models.Submission) (big.Int, error)
	GetSubmissionByID(id big.Int) (models.Submission, error)
	GetAllSubmissions() ([]models.Submission, error)
}

type Repository struct {
	Database
	Dbms
	Assign
	BannedWord
	BannedWordToAssign
	Submission
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Dbms:               NewDbmsService(db),
		Database:           NewDatabaseService(db),
		Assign:             NewAssignService(db),
		BannedWord:         NewBannedWordService(db),
		BannedWordToAssign: NewBannedWordToAssignService(db),
		Submission:         NewSubmissionService(db),
	}
}
