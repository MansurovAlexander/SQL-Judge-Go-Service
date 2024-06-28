package models

type Submission struct {
	ID           int     `json:"Id"`
	SubmissionID int     `json:"submissionId"`
	AssignID     int     `json:"assignId" binding:"required"`
	StudentID    int     `json:"studentId" binding:"required"`
	Time         float32 `json:"time"`
	Memory       float32 `json:"memory"`
	Script       string  `json:"script" binding:"required"`
	Status       int  `json:"status"`
	SubtaskID    int     `json:"subtaskId"`
}
