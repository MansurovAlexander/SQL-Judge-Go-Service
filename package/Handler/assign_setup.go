package handler

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	components "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Components"
	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	viewmodels "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/ViewModels"
	"github.com/gin-gonic/gin"
)

func (h *Handler) updateAssign(c *gin.Context) {
	var input viewmodels.AssignViewModel
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	bannedWords, admissionWords, err := components.ParseBannedWords(input.BannedWords)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	correctScripts := components.PrepareScripts(input.CorrectScript)

	databaseData, err := h.services.Database.GetDataBaseByID(input.DatabaseID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	checkScriptsResult, err := components.CheckAssignScripts(correctScripts, databaseData.Name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	for key, value := range correctScripts {
		var assignInputedData models.Assign
		key_to_int, _ := strconv.Atoi(key)

		assignInputedData.AssignID = input.AssignID
		assignInputedData.SubtaskID = key_to_int
		assignInputedData.TimeLimit = input.TimeLimit
		assignInputedData.MemoryLimit = input.MemoryLimit
		assignInputedData.CorrectScript = value
		assignInputedData.DatabaseID = input.DatabaseID
		assignInputedData.StatusID = checkScriptsResult[key_to_int].Result

		err := h.services.Assign.UpdateAssign(assignInputedData)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	err = h.services.BannedWordToAssign.DropAllBannedWordsByAssignID(input.AssignID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.services.BannedWordToAssign.CreateBannedWordToAssign(input.AssignID, bannedWords, admissionWords)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, checkScriptsResult)
}

func (h *Handler) createAssign(c *gin.Context) {
	var input viewmodels.AssignViewModel
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	bannedWords, admissionWords, err := components.ParseBannedWords(input.BannedWords)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	correctScripts := components.PrepareScripts(input.CorrectScript)

	databaseData, err := h.services.Database.GetDataBaseByID(input.DatabaseID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	checkScriptsResult, err := components.CheckAssignScripts(correctScripts, databaseData.Name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	for key, value := range correctScripts {
		var assignInputedData models.Assign
		key_to_int, _ := strconv.Atoi(key)
		assignInputedData.AssignID = input.AssignID
		assignInputedData.SubtaskID = key_to_int
		assignInputedData.TimeLimit = input.TimeLimit
		assignInputedData.MemoryLimit = input.MemoryLimit
		assignInputedData.CorrectScript = value
		assignInputedData.DatabaseID = input.DatabaseID
		assignInputedData.StatusID = checkScriptsResult[key_to_int].Result

		_, err := h.services.Assign.CreateAssign(assignInputedData)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	err = h.services.BannedWordToAssign.CreateBannedWordToAssign(input.AssignID, bannedWords, admissionWords)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, checkScriptsResult)
}

func (h *Handler) getAssignByID(c *gin.Context) {
	var input int
	input, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	assign, err := h.services.Assign.GetAssignByID(input)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, assign)
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	bannedWords, _ := h.services.BannedWordToAssign.GetBannedWordByAssignID(input)
	admissionWords, _ := h.services.BannedWordToAssign.GetAdmissionWordByAssignID(input)
	var bannedWordsForm strings.Builder
	if len(bannedWords) > len(admissionWords) {
		for key, value := range bannedWords {
			bannedWordsForm.WriteString(key)
			bannedWordsForm.WriteString(". Запрет: ")
			bannedWordsForm.WriteString(value)
			bannedWordsForm.WriteString(" ! | Допуск: ")
			bannedWordsForm.WriteString(admissionWords[key])
			bannedWordsForm.WriteString(" !;\n")
		}
	} else {
		for key, value := range admissionWords {
			bannedWordsForm.WriteString(key)
			bannedWordsForm.WriteString(". Запрет: ")
			bannedWordsForm.WriteString(bannedWords[key])
			bannedWordsForm.WriteString(" ! | Допуск: ")
			bannedWordsForm.WriteString(value)
			bannedWordsForm.WriteString(" !;\n")
		}
	}
	if len(bannedWordsForm.String()) > 0 {
		assign.BannedWords = bannedWordsForm.String()
	}
	c.JSON(http.StatusOK, assign)
}

func (h *Handler) getAllAssignes(c *gin.Context) {
	assignList, err := h.services.Assign.GetAllAssignes()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, assignList)
}

func (h *Handler) getCorrectScripts(c *gin.Context) {
	var input int
	input, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	correctScripts, err := h.services.Assign.GetAssignScriptsByID(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, correctScripts)
}

/*func (h *Handler) getAllBannedWords(c *gin.Context) {
	bannedWordList, err := h.services.BannedWord.GetAllBannedWords()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, bannedWordList)
}*/

func (h *Handler) deleteAssign(c *gin.Context) {
	var input int
	input, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.Submission.DeleteSubmissionsByAssignID(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.services.BannedWordToAssign.DropAllBannedWordsByAssignID(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.services.Assign.DeleteAssign(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{})
}
