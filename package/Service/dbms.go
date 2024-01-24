package service

import (
	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	repository "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Repository"
)

type DbmsService struct {
	repo repository.Dbms
}

func NewDbmsService(repo repository.Dbms) *DbmsService {
	return &DbmsService{repo: repo}
}

func (s *DbmsService) CreateDbms(dbms models.Dbms) (int, error){
	return s.repo.CreateDbms(dbms)
}

func (s *DbmsService) GetDbmsByID(id int) (models.Dbms, error){
	return s.repo.GetDbmsByID(id)
}

func (s *DbmsService) GetAllDbms() ([]models.Dbms, error){
	return s.repo.GetAllDbms()
}