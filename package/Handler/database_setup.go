package handler

import (
	"net/http"
	"strconv"

	models "github.com/MansurovAlexander/SQL-Judge-Moodle-Plugin/package/Models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createDataBase(c *gin.Context) {
	var input models.Database
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Database.CreateDatabase(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) createDbms(c *gin.Context) {
	var input models.Dbms
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Dbms.CreateDbms(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getDbmsByID(c *gin.Context) {
	input, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	dbms, err := h.services.Dbms.GetDbmsByID(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, dbms)
}

func (h *Handler) getAllDbms(c *gin.Context) {
	dbmsList, err := h.services.Dbms.GetAllDbms()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, dbmsList)
}

func (h *Handler) getDataBaseByID(c *gin.Context) {
	input, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	database, err := h.services.Database.GetDataBaseByID(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, database)
}

func (h *Handler) getAllDataBases(c *gin.Context) {
	databasesList, err := h.services.Database.GetAllDatabases()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, databasesList)
}
