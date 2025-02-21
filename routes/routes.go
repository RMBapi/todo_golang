package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/api/v1/todo", HandleRequest)
	server.GET("/api/v1/todo", HandleRequest)
	server.PUT("/api/v1/update-todo/:id", UpdateTask)
}
