package models

type BannedWordToAssign struct {
	ID           int `json:"id"`
	AssignID     int `json:"assign" binding:"required"`
	BannedWordID int `json:"banned" binding:"required"`
}
