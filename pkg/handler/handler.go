package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"ims-staff-api/pkg/service"

	_ "ims-staff-api/docs"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	task := router.Group("/task") //h.userIdentity
	{
		task.GET("", h.GetTasksList)                  // Получение всех задач
		task.GET("/status/:name", h.GetTasksByStatus) // Получение списка задач по статусу
		task.GET("/:id", h.GetTaskById)               //Получение задачи по ID

		task.POST("", h.CreateTask) // Создать задачу

		task.PUT("/:id", h.UpdateTask) // Редактировать задачу по ID

		task.DELETE("/:id", h.DeleteTask) // Удалить задачу

		status := task.Group("/status")
		{
			status.GET("", h.GetStatusList) // Получение всех статусов задач
		}
	}

	management := router.Group("/management") //h.userIdentity
	{
		management.GET("", h.GetStaffList)           // Посмотреть всех сотрудиков и их задачи
		management.GET("/:id", h.GetWorkerTasksById) // Посмотреть список задач сотрудника по ID

		management.POST("/:worker_id/task/:task_id", h.AssignTask) // Менеджер отдает задачу для работника по его ID
	}

	return router
}
