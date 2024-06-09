package models

type Database struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	FileName    string `json:"fileName" binding:"required"`
	Extension   string `json:"fileExtension" binding:"required"`
	Base64File  string `json:"base64File" binding:"required"`
	Description string `json:"description"`
	DbmsID      string `json:"dbms" binding:"required"`
}
