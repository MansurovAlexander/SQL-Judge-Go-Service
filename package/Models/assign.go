package models

import "math/big"

type Assign struct {
	ID            big.Int `json:"id" binding:"required"`
	TimeLimit     int     `json:"time"`
	MemoryLimit   int     `json:"memory"`
	CorrectScript string  `json:"script" binding:"required"`
	DatabaseID    int     `json:"database" binding:"required"`
}
