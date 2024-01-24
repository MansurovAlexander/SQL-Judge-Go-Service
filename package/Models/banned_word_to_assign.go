package models

import "math/big"

type BannedWordToAssign struct {
	ID           int     `json:"id"`
	AssignID     big.Int `json:"assign" binding:"required"`
	BannedWordID int     `json:"banned" binding:"required"`
}
