package service

import (
	"math/big"

	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	repository "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Repository"
)

type AssignService struct {
	repo repository.Assign
}

func NewAssignService(repo repository.Assign) *AssignService {
	return &AssignService{repo}
}

func (s *AssignService) CreateAssign(assign models.Assign) (big.Int, error) {
	return s.repo.CreateAssign(assign)
}

func (s *AssignService) GetAssignByID(id big.Int) (models.Assign, error) {
	return s.repo.GetAssignByID(id)
}

func (s *AssignService) GetAllAssignes() ([]models.Assign, error) {
	return s.repo.GetAllAssignes()
}
