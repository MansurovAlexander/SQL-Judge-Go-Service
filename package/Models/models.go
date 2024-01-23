package models

type Submission struct {
	ID             int
	MoodleAssignID int
	StudentID      int
	Time           int
	Memory         int
	Grade          string
	Script         string
}

type DataBases struct {
	ID             int
	DbmsID         int
	Name           string
	Description    string
	CreationScript string
}

type Assign struct {
	MoodleAssignID int
	Name           string
	Description    string
	TimeLimit      int
	MemoryLimit    int
	CorrectScript  string
	BannedWords    string
	DbmsID         int
}

type Dbms struct {
	ID   int
	Name string
}
