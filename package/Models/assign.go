package models

type Assign struct {
	ID            int     `json:"id"`
	AssignID      int     `json:"assignId" binding:"required"`
	SubtaskID     int     `json:"subtaskId"`
	TimeLimit     float32 `json:"time"`
	MemoryLimit   float32 `json:"memory"`
	CorrectScript string  `json:"script" binding:"required"`
	DatabaseID    int     `json:"database" binding:"required"`
	StatusID      int     `json:"statusId`
}
