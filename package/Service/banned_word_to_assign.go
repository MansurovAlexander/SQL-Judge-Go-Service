package service

import (
	repository "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Repository"
)

type BannedWordToAssignService struct {
	repo repository.BannedWordToAssign
}

func NewBannedWordToAssignService(repo repository.BannedWordToAssign) *BannedWordToAssignService {
	return &BannedWordToAssignService{repo}
}

func (s *BannedWordToAssignService) CreateBannedWordToAssign(assignID int, bannedWordID int) (int, error) {
	return s.repo.CreateBannedWordToAssign(assignID, bannedWordID)
}

func (s *BannedWordToAssignService) GetBannedWordByAssignID(id int) ([]int, error) {
	return s.repo.GetBannedWordByAssignID(id)
}
