package service

import repository "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Repository"

type Submission interface {
}
type DataBase interface {
}

type Service struct {
	Submission
	DataBase
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
