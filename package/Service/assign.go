package service

import (
	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	repository "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Repository"
	viewmodels "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/ViewModels"
)

type AssignService struct {
	repo repository.Assign
}

func NewAssignService(repo repository.Assign) *AssignService {
	return &AssignService{repo}
}

func (s *AssignService) CreateAssign(assign models.Assign) (int, error) {
	return s.repo.CreateAssign(assign)
}

func (s *AssignService) UpdateAssign(assign models.Assign) error {
	return s.repo.UpdateAssign(assign)
}

func (s *AssignService) GetAssignByID(id int) (viewmodels.AssignViewModel, error) {
	return s.repo.GetAssignByID(id)
}

func (s *AssignService) GetAllAssignes() ([]models.Assign, error) {
	return s.repo.GetAllAssignes()
}

func (s *AssignService) DeleteAssign(id int) error {
	return s.repo.DeleteAssign(id)
}

func (s *AssignService) GetAssignScriptsByID(id int) (map[int]models.Assign, error) {
	return s.repo.GetAssignScriptsByID(id)
}
