package viewmodels

type AssignViewModel struct {
	ID            int    `json:"id"`
	AssignID      int    `json:"assignId" binding:"required"`
	SubtaskID     int    `json:"subtaskId"`
	TimeLimit     int    `json:"time"`
	MemoryLimit   int    `json:"memory"`
	CorrectScript string `json:"script" binding:"required"`
	DatabaseID    int    `json:"database" binding:"required"`
	BannedWords   string `json:"bannedwords"`
}
