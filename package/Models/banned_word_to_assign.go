package models

import "math/big"

type BannedWordToAssign struct {
	ID           int
	AssignID     big.Int
	BannedWordID int
}
