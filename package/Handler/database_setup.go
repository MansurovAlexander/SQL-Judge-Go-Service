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
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":   dbms.ID,
		"name": dbms.Name,
	})
}

func (h *Handler) getAllDbms(c *gin.Context) {

}

func (h *Handler) getDataBaseByID(c *gin.Context) {

}

func (h *Handler) getAllDataBases(c *gin.Context) {

}
