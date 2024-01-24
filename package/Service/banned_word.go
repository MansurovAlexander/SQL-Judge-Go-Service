package service

import (
	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	repository "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Repository"
)

type BannedWordService struct {
	repo repository.BannedWord
}

func NewBannedWordService(repo repository.BannedWord) *BannedWordService {
	return &BannedWordService{repo: repo}
}

func (s *BannedWordService) CreateBannedWord(bannedWord models.BannedWord) (int, error) {
	return s.repo.CreateBannedWord(bannedWord)
}

func (s *BannedWordService) GetBannedWordByID(id int) (models.BannedWord, error) {
	return s.repo.GetBannedWordByID(id)
}

func (s *BannedWordService) GetAllBannedWords() ([]models.BannedWord, error) {
	return s.repo.GetAllBannedWords()
}
