package service

import (
	"math/big"

	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	repository "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Repository"
)

type SubmissionService struct {
	repo repository.Submission
}

func NewSubmissionService(repo repository.Submission) *SubmissionService {
	return &SubmissionService{repo}
}

func (s *SubmissionService) CreateSubmission(submission models.Submission) (big.Int, error) {
	return s.repo.CreateSubmission(submission)
}

func (s *SubmissionService) GetSubmissionByID(id big.Int) (models.Submission, error) {
	return s.repo.GetSubmissionByID(id)
}

func (s *SubmissionService) GetAllSubmissions() ([]models.Submission, error) {
	return s.repo.GetAllSubmissions()
}
