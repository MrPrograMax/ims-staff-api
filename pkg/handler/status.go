package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetStatusList(c *gin.Context) {
	statuses, err := h.services.Status.GetStatusList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error of get statuses")
		return
	}

	c.JSON(http.StatusOK, statuses)
}
