package handler

import (
	"math/big"
	"net/http"

	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	viewmodels "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/ViewModels"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createBannedWord(c *gin.Context) {
	var input models.BannedWord
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.BannedWord.CreateBannedWord(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) createAssign(c *gin.Context) {
	var input viewmodels.AssignViewModel
	var bannedWordsToAssignIds []big.Int
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var assignInputedData models.Assign
	assignInputedData.TimeLimit = input.TimeLimit
	assignInputedData.MemoryLimit = input.MemoryLimit
	assignInputedData.CorrectScript = input.CorrectScript
	assignInputedData.DatabaseID = input.DatabaseID
	bannedWords := input.BannedWords
	id, err := h.services.Assign.CreateAssign(assignInputedData)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	for i := range bannedWords {
		bwToAssignId, err := h.services.BannedWordToAssign.CreateBannedWordToAssign(id, bannedWords[i])
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		bannedWordsToAssignIds = append(bannedWordsToAssignIds, bwToAssignId)
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAssignByID(c *gin.Context) {
	var input big.Int
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	assign, err := h.services.Assign.GetAssignByID(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
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

func (h *Handler) getAllBannedWords(c *gin.Context) {
	bannedWordList, err := h.services.BannedWord.GetAllBannedWords()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, bannedWordList)
}
