package router

import (
	"task_manager/controllers"
	"task_manager/data"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	taskService := data.NewTaskService()
	taskController := controllers.NewTaskController(taskService)

	api := r.Group("/api")
	{
		tasks := api.Group("/tasks")
		{
			tasks.GET("", taskController.GetAllTasks)
			tasks.POST("", taskController.CreateTask)
			tasks.GET("/:id", taskController.GetTaskByID)
			tasks.PUT("/:id", taskController.UpdateTask)
			tasks.DELETE("/:id", taskController.DeleteTask)
		}
	}

	return r
}