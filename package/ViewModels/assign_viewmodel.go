package viewmodels

import "math/big"

type AssignViewModel struct {
	ID            big.Int `json:"id" binding:"required"`
	TimeLimit     int     `json:"time"`
	MemoryLimit   int     `json:"memory"`
	CorrectScript string  `json:"script" binding:"required"`
	DatabaseID    int     `json:"database" binding:"required"`
	BannedWords   []int   `json:"bannedwords"`
}
