package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	components "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Components"
	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	"github.com/gin-gonic/gin"
)

// Функция проверки sql запроса, сразу после загрузки на сервер
// SQL query verification function, immediately after uploading to the server
// Запрос приходит в формате JSON, содержащем Id задания, Id студента и скрипт
func (h *Handler) checkSubmission(c *gin.Context) {
	var input models.Submission
	//парсим json данные по модели submission
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
	//получаем запрещенные слова
	var bannedWords, admissionWords map[string]string
	bannedWords, err = h.services.GetBannedWordByAssignID(input.AssignID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	admissionWords, err = h.services.GetAdmissionWordByAssignID(input.AssignID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	//подготавливаем скрипт для выполнения
	studentScripts := components.PrepareScripts(input.Script)
	correctScripts := components.PrepareScripts(assignData.CorrectScript)

	id, err := h.services.Judge.CheckSubmission(studentScripts, correctScripts, bannedWords, admissionWords, databaseData.Name, input.SubmissionID, input.AssignID, input.StudentID, assignData.TimeLimit, assignData.MemoryLimit)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllSubmissions(c *gin.Context) {

}

func (h *Handler) getSubmissionByID(c *gin.Context) {
	studentId, err := strconv.Atoi(c.Param("student_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	assignId, err := strconv.Atoi(c.Param("assign_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	submissionList, err := h.services.Submission.GetSubmissionByID(studentId, assignId)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, submissionList)
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, submissionList)
}
