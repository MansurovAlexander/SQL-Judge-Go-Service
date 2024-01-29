package service

import (
	"math/big"

	repository "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Repository"
)

type BannedWordToAssignService struct {
	repo repository.BannedWordToAssign
}

func NewBannedWordToAssignService(repo repository.BannedWordToAssign) *BannedWordToAssignService {
	return &BannedWordToAssignService{repo}
}

func (s *BannedWordToAssignService) CreateBannedWordToAssign(assignID big.Int, bannedWordID int) (big.Int, error) {
	return s.repo.CreateBannedWordToAssign(assignID, bannedWordID)
}

func (s *BannedWordToAssignService) GetBannedWordByAssignID(id big.Int) ([]int, error) {
	return s.repo.GetBannedWordByAssignID(id)
}
