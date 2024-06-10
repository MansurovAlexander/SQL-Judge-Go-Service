package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllStatuses(c *gin.Context) {
	statuses, err := h.services.Status.GetAllStatuses()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statuses)
}
