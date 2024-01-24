package models

type BannedWord struct {
	ID         int    `json:"id"`
	BannedWord string `json:"banned" binding:"required"`
}
