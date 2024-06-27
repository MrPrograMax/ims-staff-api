package handler

import (
	"github.com/gin-gonic/gin"
	"ims-staff-api/models"
	"net/http"
	"strconv"
)

func (h *Handler) GetTasksList(c *gin.Context) {
	tasks, err := h.services.GetTasksList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error of get all tasks")
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) GetTasksByStatus(c *gin.Context) {
	name := c.Param("name")

	tasks, err := h.services.GetTasksByStatus(name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error of get task by ID")
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) CreateTask(c *gin.Context) {
	var input models.Task
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	statusId, err := h.services.GetStatusIdByName("Ожидает")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error of get status name")
		return
	}

	id, err := h.services.CreateTask(input, statusId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error of create")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) UpdateTask(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID param")
		return
	}

	var input models.Task
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := h.services.UpdateTask(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error of update group")
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})
}

func (h *Handler) DeleteTask(c *gin.Context) {
	taskId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID param")
		return
	}

	statusId, err := h.services.GetStatusIdByName("Отменена")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error of create")
		return
	}

	if err := h.services.DeleteTask(taskId, statusId); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error of delete group")
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})
}
