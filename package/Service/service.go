package service

import (
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
	CreateAssign(assign models.Assign) (int, error)
	GetAssignByID(id int) (models.Assign, error)
	GetAllAssignes() ([]models.Assign, error)
}

type BannedWordToAssign interface {
	CreateBannedWordToAssign(assignID int, bannedWordID int) (int, error)
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

type Service struct {
	Database
	Dbms
	BannedWord
	Assign
	BannedWordToAssign
	Submission
	Judge
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Dbms:               NewDbmsService(repos.Dbms),
		Database:           NewDatabaseService(repos.Database),
		Assign:             NewAssignService(repos.Assign),
		BannedWord:         NewBannedWordService(repos.BannedWord),
		BannedWordToAssign: NewBannedWordToAssignService(repos.BannedWordToAssign),
		Submission:         NewSubmissionService(repos.Submission),
		Judge:              NewJudgeService(repos.Judge),
	}
}
