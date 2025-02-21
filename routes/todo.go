package routes

import (
	"net/http"

	"example.com/todo/models"
	"github.com/gin-gonic/gin"
)

func HandleRequest(context *gin.Context) {
	if context.Request.Method == http.MethodPost {
		var task models.TodoList

		type requestBody struct {
			Task        *string `json:"task"`
			Description *string `json:"description"`
			Datetime    *string `json:"datetime"`
		}

		var req requestBody
		err := context.ShouldBindJSON(&req)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Could not purse request"})
			return
		}

		task.Task = *req.Task
		task.Description = *req.Description
		task.Datetime = *req.Datetime

		err = task.Save()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not insert item"})
			return
		}

	} else if context.Request.Method == http.MethodGet {

		var task models.TodoList

		taskList, err := task.ViewTask()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find task successfully"})
			return
		}

		context.Header("Content-Type", "application/json")
		context.JSON(http.StatusOK, taskList)
	}
}
