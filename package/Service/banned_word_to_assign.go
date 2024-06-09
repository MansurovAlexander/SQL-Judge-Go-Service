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

func (s *BannedWordToAssignService) CreateBannedWordToAssign(assignID int, bannedWords, admissionWords map[string][]string) error {
	return s.repo.CreateBannedWordToAssign(assignID, bannedWords, admissionWords)
}

func (s *BannedWordToAssignService) GetBannedWordByAssignID(id int) (map[string]string, error) {
	return s.repo.GetBannedWordByAssignID(id)
}

func (s *BannedWordToAssignService) GetAdmissionWordByAssignID(id int) (map[string]string, error) {
	return s.repo.GetAdmissionWordByAssignID(id)
}

func (s *BannedWordToAssignService) DropAllBannedWordsByAssignID(assignId int) error {
	return s.repo.DropAllBannedWordsByAssignID(assignId)
}
