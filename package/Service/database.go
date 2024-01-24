package service

import (
	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	repository "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Repository"
)

type DatabaseService struct {
	repo repository.Database
}

func NewDatabaseService(repo repository.Database) *DatabaseService {
	return &DatabaseService{repo: repo}
}

func (s *DatabaseService) CreateDatabase(database models.Database) (int, error) {
	return s.repo.CreateDatabase(database)
}

func (s *DatabaseService) GetDataBaseByID(id int) (models.Database, error) {
	return s.repo.GetDatabaseByID(id)
}

func (s *DatabaseService) GetAllDatabases() ([]models.Database, error) {
	return s.repo.GetAllDatabases()
}
