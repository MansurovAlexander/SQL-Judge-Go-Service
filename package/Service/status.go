package service

import (
	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	repository "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Repository"
)

type StatusService struct {
	repo repository.Status
}

func NewStatusService(repo repository.Status) *StatusService {
	return &StatusService{repo}
}

func (s *StatusService) GetAllStatuses() ([]models.Status, error) {
	return s.repo.GetAllStatuses()
}
