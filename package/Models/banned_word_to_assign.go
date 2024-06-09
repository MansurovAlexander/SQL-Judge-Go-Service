package models

type BannedWordToAssign struct {
	ID             int    `json:"id"`
	AssignID       int    `json:"assign" binding:"required"`
	BannedWords    string `json:"banned" `
	AdmissionWords string `json:"admission"`
	SubtaskID      int    `json:"subtaskId"`
}
