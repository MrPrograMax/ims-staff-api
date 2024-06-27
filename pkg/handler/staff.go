package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetStaffList(c *gin.Context) {
	staff, err := h.services.GetStaffList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error of get all groups")
		return
	}

	c.JSON(http.StatusOK, staff)
}

func (h *Handler) GetWorkerTasksById(c *gin.Context) {
	workerId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID param")
		return
	}

	tasks, err := h.services.GetWorkerTasksById(workerId)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) AssignTask(c *gin.Context) {
	workerId, err := strconv.ParseInt(c.Param("worker_id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID param")
		return
	}

	taskId, err := strconv.ParseInt(c.Param("task_id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID param")
		return
	}

	task, err := h.services.GetTaskByID(taskId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error of assign task")
	}

	task.UserId = workerId

	if err := h.services.UpdateTask(taskId, task); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})
}
