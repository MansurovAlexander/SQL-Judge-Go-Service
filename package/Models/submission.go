package models

type Submission struct {
	ID        int    `json:"submissionId"`
	AssignID  int    `json:"assignId" binding:"required"`
	StudentID int    `json:"studentId" binding:"required"`
	Time      int    `json:"time"`
	Memory    int    `json:"memory"`
	Script    string `json:"script" binding:"required"`
	Grade     string `json:"grade"`
}
