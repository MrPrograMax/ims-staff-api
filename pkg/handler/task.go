package handler

import (
	"github.com/gin-gonic/gin"
	"ims-staff-api/models"
	"net/http"
	"strconv"
)

// GetTasksList godoc
// @Summary Get list of tasks
// @Description Get list of all tasks
// @Tags Task
// @Accept json
// @Produce json
// @Success 200 {array} models.Task
// @Failure 204 {object} errorResponse
// @Router /task [get]
func (h *Handler) GetTasksList(c *gin.Context) {
	tasks, err := h.services.GetTasksList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error of get all tasks")
		return
	}

	if tasks == nil {
		c.JSON(http.StatusNoContent, tasks)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// GetTasksByStatus godoc
// @Summary Get tasks by status
// @Description Get list of tasks by status name
// @Tags Status
// @Accept json
// @Produce json
// @Param name path string true "Status Name"
// @Success 200 {array} models.Task
// @Failure 204 {object} errorResponse
// @Router /task/status/{name} [get]
func (h *Handler) GetTasksByStatus(c *gin.Context) {
	name := c.Param("name")

	tasks, err := h.services.GetTasksByStatus(name)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error of get task by ID")
		return
	}

	if tasks == nil {
		c.JSON(http.StatusNoContent, tasks)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// CreateTask godoc
// @Summary Create a new task
// @Description Create a new task with the provided information
// @Tags Task
// @Accept json
// @Produce json
// @Param input body models.Task true "Task info"
// @Success 200 {object} int
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /task [post]
func (h *Handler) CreateTask(c *gin.Context) {
	var input models.Task
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	statusId, err := h.services.GetStatusIdByName("Waiting")
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

// GetTaskById godoc
// @Summary Get task by ID
// @Description Get task details by ID
// @Tags Task
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} models.Task
// @Failure 400 {object} errorResponse
// @Failure 204 {object} errorResponse
// @Router /task/{id} [get]
func (h *Handler) GetTaskById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID param")
		return
	}

	task, err := h.services.GetTaskByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusNoContent, "Error of get task by ID")
		return
	}

	c.JSON(http.StatusOK, task)
}

// UpdateTask godoc
// @Summary Update task
// @Description Update task details by ID
// @Tags Task
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param input body models.UpdateTask true "Update Task info"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /task/{id} [put]
func (h *Handler) UpdateTask(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID param")
		return
	}

	var input models.UpdateTask
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := h.services.UpdateTask(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})
}

// DeleteTask godoc
// @Summary Delete task
// @Description Delete task by ID
// @Tags Task
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {object} statusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /task/{id} [delete]
func (h *Handler) DeleteTask(c *gin.Context) {
	taskId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid ID param")
		return
	}

	statusId, err := h.services.GetStatusIdByName("Canceled")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Error of cancel status")
		return
	}

	task, err := h.services.GetTaskByID(taskId)
	if err != nil {
		newErrorResponse(c, http.StatusNoContent, "Error of get task by ID")
		return
	}
	task.StatusId = statusId

	updateTask := models.UpdateTask{
		StatusId: &statusId,
	}

	err = h.services.UpdateTask(taskId, updateTask)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error()) //"Error of update task"
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "OK",
	})
}
