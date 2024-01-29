package handler

import (
	"net/http"

	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	"github.com/gin-gonic/gin"
)

// Функция проверки sql запроса, сразу после загрузки на сервер
// SQL query verification function, immediately after uploading to the server
func (h *Handler) checkSubmission(c *gin.Context) {
	var input models.Submission
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

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
