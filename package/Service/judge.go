package service

import (
	repository "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Repository"
)

type JudgeService struct {
	repo repository.Judge
}

func NewJudgeService(repo repository.Judge) *JudgeService {
	return &JudgeService{repo}
}

func (s *JudgeService) CheckSubmission(inputedScript, correctScript, banned_words, admission_words map[string]string, dbFileName string, submissionID, assignID, studentID, timeLimit, memoryLimit int) (int, error) {
	return s.repo.CheckSubmission(inputedScript, correctScript, banned_words, admission_words, dbFileName, submissionID, assignID, studentID, timeLimit, memoryLimit)
}
