package viewmodels

type DatabaseViewModel struct {
	ID          int    `json:"id" binding:"required"`
	FileName    string `json:"fileName"`
	Name        string `json:"dbName"`
	Description string `json:"description"`
	DbmsName    string `json:"dbmsName" binding:"required"`
}
