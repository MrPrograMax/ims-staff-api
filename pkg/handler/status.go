package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetStatusList godoc
// @Summary Get list of statuses
// @Description Get list of all task statuses
// @Tags Status
// @Accept json
// @Produce json
// @Success 200 {array} models.TaskStatus
// @Failure 204 {object} errorResponse
// @Router /task/status [get]
func (h *Handler) GetStatusList(c *gin.Context) {
	statuses, err := h.services.Status.GetStatusList()
	if err != nil {
		newErrorResponse(c, http.StatusNoContent, "Error of get statuses")
		return
	}

	c.JSON(http.StatusOK, statuses)
}
