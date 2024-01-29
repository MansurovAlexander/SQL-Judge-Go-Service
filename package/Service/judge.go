package service

import (
	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	repository "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Repository"
)

type JudgeService struct {
	repo repository.Judge
}

func NewJudgeService(repo repository.Judge) *JudgeService {
	return &JudgeService{repo}
}

func (s *JudgeService) CheckSubmisson(submission models.Submission) (string, error) {
	return s.repo.CheckSubmission(submission)
}
