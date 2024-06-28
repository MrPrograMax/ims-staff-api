package handler

import (
	"github.com/gin-gonic/gin"
	"ims-staff-api/models"
	"net/http"
	"strconv"
)

// GetStaffList godoc
// @Summary Get list of staff
// @Description Get list of all staff members
// @Tags Worker
// @Accept json
// @Produce json
// @Success 200 {array} models.UserDto
// @Failure 500 {object} errorResponse
// @Router /management [get]
func (h *Handler) GetStaffList(c *gin.Context) {
	staff, err := h.services.GetStaffList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error of get all staff")
		return
	}

	if staff == nil {
		c.JSON(http.StatusNoContent, staff)
		return
	}

	c.JSON(http.StatusOK, staff)
}

// GetWorkerTasksById godoc
// @Summary Get worker's tasks by ID
// @Description Get list of tasks assigned to a worker by worker ID
// @Tags Worker
// @Accept json
// @Produce json
// @Param id path int true "Worker ID"
// @Success 200 {array} models.Task
// @Failure 400 {object} errorResponse
// @Router /management/{id} [get]
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

// AssignTask godoc
// @Summary Assign a task to a worker
// @Description Assign a task to a worker by their IDs
// @Tags Worker
// @Accept json
// @Produce json
// @Param worker_id path int true "Worker ID"
// @Param task_id path int true "Task ID"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /management/{worker_id}/task/{task_id} [post]
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

	statusId, err := h.services.GetStatusIdByName("Waiting")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error of get status name")
		return
	}

	updateTask := models.UpdateTask{
		Title:       &task.Title,
		Description: &task.Description,
		StatusId:    &statusId,
		UserId:      &workerId,
	}

	if err := h.services.UpdateTask(taskId, updateTask); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})
}
