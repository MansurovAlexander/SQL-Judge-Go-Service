package models

type Dbms struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}
