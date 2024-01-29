package service

import (
	"math/big"

	repository "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Repository"
)

type JudgeService struct {
	repo repository.Judge
}

func NewJudgeService(repo repository.Judge) *JudgeService {
	return &JudgeService{repo}
}

func (s *JudgeService) CheckSubmission(inputedScript, dbCreateScript, correctScript, dbName string, assignID, studentID big.Int, timeLimit, memoryLimit int) (big.Int, error) {
	return s.repo.CheckSubmission(inputedScript, dbCreateScript, correctScript, dbName, assignID, studentID, timeLimit, memoryLimit)
}
