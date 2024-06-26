package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetStaffList(c *gin.Context) {
	c.JSON(http.StatusOK, statusResponse{
		Status: "TODO GetStaffList",
	})
}

func (h *Handler) GetWorkerTasksById(c *gin.Context) {
	c.JSON(http.StatusOK, statusResponse{
		Status: "TODO GetWorkerTasksById",
	})
}

func (h *Handler) AssignTask(c *gin.Context) {
	c.JSON(http.StatusOK, statusResponse{
		Status: "TODO AssignTask",
	})
}
