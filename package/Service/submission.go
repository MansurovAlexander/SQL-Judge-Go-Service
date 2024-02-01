package service

import (
	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	repository "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Repository"
)

type SubmissionService struct {
	repo repository.Submission
}

func NewSubmissionService(repo repository.Submission) *SubmissionService {
	return &SubmissionService{repo}
}

func (s *SubmissionService) CreateSubmission(submission models.Submission) (int, error) {
	return s.repo.CreateSubmission(submission)
}

func (s *SubmissionService) GetSubmissionByID(id int) (models.Submission, error) {
	return s.repo.GetSubmissionByID(id)
}

func (s *SubmissionService) GetAllSubmissions() ([]models.Submission, error) {
	return s.repo.GetAllSubmissions()
}
