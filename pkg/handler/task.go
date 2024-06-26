package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetTasksList(c *gin.Context) {
	c.JSON(http.StatusOK, statusResponse{
		Status: "TODO GetTasksList",
	})
}

func (h *Handler) GetTasksByStatus(c *gin.Context) {
	c.JSON(http.StatusOK, statusResponse{
		Status: "TODO GetTasksByStatus",
	})
}

func (h *Handler) CreateTask(c *gin.Context) {
	c.JSON(http.StatusOK, statusResponse{
		Status: "TODO CreateTask",
	})
}

func (h *Handler) UpdateTask(c *gin.Context) {
	c.JSON(http.StatusOK, statusResponse{
		Status: "TODO UpdateTask",
	})
}

func (h *Handler) DeleteTask(c *gin.Context) {
	c.JSON(http.StatusOK, statusResponse{
		Status: "TODO DeleteTask",
	})
}
