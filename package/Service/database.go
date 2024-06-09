package service

import (
	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	repository "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Repository"
	viewmodels "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/ViewModels"
)

type DatabaseService struct {
	repo repository.Database
}

func NewDatabaseService(repo repository.Database) *DatabaseService {
	return &DatabaseService{repo: repo}
}

func (s *DatabaseService) DeleteDataBaseByID(id int) (string, error) {
	return s.repo.DeleteDataBaseByID(id)
}

func (s *DatabaseService) CreateDatabase(database models.Database) (int, error) {
	return s.repo.CreateDatabase(database)
}

func (s *DatabaseService) GetDataBaseByID(id int) (viewmodels.DatabaseViewModel, error) {
	return s.repo.GetDatabaseByID(id)
}

func (s *DatabaseService) GetAllDatabases() ([]viewmodels.DatabaseViewModel, error) {
	return s.repo.GetAllDatabases()
}
