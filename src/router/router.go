package router

import (
	"GolangTodoApp/src/controller"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.New()

	router.GET("/", controller.ListTodo)
	router.GET("/:id", controller.GetTodo)
	router.POST("/", controller.AddTodo)
	router.POST("/update/:id", controller.UpdateTodo)
	router.POST("/remove/:id", controller.DeleteTodo)

	router.Run(":8888")
}
