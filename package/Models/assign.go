package models

import "math/big"

type Assign struct {
	ID            big.Int
	TimeLimit     int
	MemoryLimit   int
	CorrectScript string
	DatabaseID    int
}
