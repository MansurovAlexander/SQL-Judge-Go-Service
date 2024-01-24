package service

import (
	"math/big"

	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	repository "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Repository"
)

type Database interface {
	CreateDatabase(database models.Database) (int, error)
	GetAllDatabases() ([]models.Database, error)
	GetDataBaseByID(id int) (models.Database, error)
}

type Dbms interface {
	CreateDbms(dbms models.Dbms) (int, error)
	GetAllDbms() ([]models.Dbms, error)
	GetDbmsByID(id int) (models.Dbms, error)
}

type BannedWord interface {
	CreateBannedWord(bannedWord models.BannedWord) (int, error)
	GetAllBannedWords() ([]models.BannedWord, error)
}

type Assign interface {
	CreateAssign(assign models.Assign) (big.Int, error)
	GetAssignByID(id big.Int) (models.Assign, error)
	GetAllAssignes() ([]models.Assign, error)
}

type BannedWordToAssign interface {
	CreateBannedWordToAssign(bannedWordToAssign models.BannedWordToAssign) (big.Int, error)
	GetBannedWordByAssignID(assignId big.Int) ([]int, error)
}

type Submission interface {
	CreateSubmission(submission models.Submission) (big.Int, error)
	GetSubmissionByID(id big.Int) (models.Submission, error)
	GetAllSubmissions() ([]models.Submission, error)
}

type Service struct {
	Database
	Dbms
	BannedWord
	Assign
	BannedWordToAssign
	Submission
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Dbms:               NewDbmsService(repos.Dbms),
		Database:           NewDatabaseService(repos.Database),
		Assign:             NewAssignService(repos.Assign),
		BannedWord:         NewBannedWordService(repos.BannedWord),
		BannedWordToAssign: NewBannedWordToAssignService(repos.BannedWordToAssign),
		Submission:         NewSubmissionService(repos.Submission),
	}
}
