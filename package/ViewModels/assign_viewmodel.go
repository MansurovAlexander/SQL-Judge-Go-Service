package viewmodels

type AssignViewModel struct {
	ID            int    `json:"id" binding:"required"`
	TimeLimit     int    `json:"time"`
	MemoryLimit   int    `json:"memory"`
	CorrectScript string `json:"script" binding:"required"`
	DatabaseID    int    `json:"database" binding:"required"`
	BannedWords   []int  `json:"bannedwords"`
}
