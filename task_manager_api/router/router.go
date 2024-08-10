package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.POST("/register", controllers.SignUp)
	router.POST("/login", controllers.LogIn)
	// only for logged in users 
	router.GET("/tasks",middleware.UserMiddleware, controllers.GetTasks)
	router.GET("/tasks/:id",middleware.UserMiddleware, controllers.GetTask)
	// only for Admin
	router.PUT("/promote/:email",middleware.AdminMiddleWare, controllers.PromoteUser)
	router.POST("/tasks",middleware.AdminMiddleWare, controllers.CreateTask)
	router.PUT("/tasks/:id",middleware.AdminMiddleWare, controllers.UpdateTask)
	router.DELETE("/tasks/:id",middleware.AdminMiddleWare, controllers.DeleteTask)

	router.Run()
}
