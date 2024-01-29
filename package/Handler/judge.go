package handler

import (
	"net/http"

	methods "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Methods"
	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	"github.com/gin-gonic/gin"
)

// Функция проверки sql запроса, сразу после загрузки на сервер
// SQL query verification function, immediately after uploading to the server
// Запрос приходит в формате JSON, содержащем Id задания, Id студента и скрипт
func (h *Handler) checkSubmission(c *gin.Context) {
	var input models.Submission
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	assignData, err := h.services.Assign.GetAssignByID(input.AssignID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	databaseData, err := h.services.Database.GetDataBaseByID(assignData.DatabaseID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	dbmsData, err := h.services.Dbms.GetDbmsByID(databaseData.DbmsID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	databaseData.CreationScript, err = methods.PrepareDBScript(databaseData.CreationScript, dbmsData.Name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id, err := h.services.Judge.CheckSubmission(input.Script, databaseData.CreationScript, assignData.CorrectScript, dbmsData.Name, input.AssignID, input.StudentID, assignData.TimeLimit, assignData.MemoryLimit)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
	//После обработки sql запроса на инъекции создается запись в бд с попыткой студента
	/*id, err := h.services.Submission.CreateSubmission(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}*.
	/*c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})*/
}

func (h *Handler) getAllSubmissions(c *gin.Context) {

}

func (h *Handler) getSubmissionByID(c *gin.Context) {

}
