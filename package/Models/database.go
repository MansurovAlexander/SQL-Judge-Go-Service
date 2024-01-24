package models

type Database struct {
	ID             int    `json:"id"`
	Name           string `json:"name" binding:"required"`
	Description    string `json:"description"`
	CreationScript string `json:"script" binding:"required"`
	DbmsID         int    `json:"dbms" binding:"required"`
}
