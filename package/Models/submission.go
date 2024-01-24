package models

import "math/big"

type Submission struct {
	ID        big.Int `json:"submissionId" binding:"required"` 
	AssignID  big.Int `json:"assignId" binding:"required"`
	StudentID big.Int `json:"studentId" binding:"required"`
	Time      int `json:"time"`
	Memory    int `json:"memory"`
	Script    string `json:"script" binding:"required"`
	Grade     string `json:"grade"`
}
