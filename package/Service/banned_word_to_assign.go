package service

import (
	"math/big"

	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	repository "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Repository"
)

type BannedWordToAssignService struct {
	repo repository.BannedWordToAssign
}

func NewBannedWordToAssignService(repo repository.BannedWordToAssign) *BannedWordToAssignService {
	return &BannedWordToAssignService{repo}
}

func (s *BannedWordToAssignService) CreateBannedWordToAssign(bannedWordToAssign models.BannedWordToAssign) (big.Int, error) {
	return s.repo.CreateBannedWordToAssign(bannedWordToAssign)
}

func (s *BannedWordToAssignService) GetBannedWordByAssignID(id big.Int) ([]int, error) {
	return s.repo.GetBannedWordByAssignID(id)
}
